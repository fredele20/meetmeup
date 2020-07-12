package graph

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"meetmeup/graph/generated"
	"meetmeup/graph/model"
	"meetmeup/middleware"
	"strconv"
	"time"
)

var (
	ErrBadCredentials  = errors.New("email/password combination does not match")
	ErrUnAuthenticated = errors.New("unauthenticated")
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := m.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("user with email already exists")
	}

	_, err = m.UserRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("user with username already exists")
	}
	payload := &model.User{
		ID:        strconv.Itoa(rand.Int()),
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = payload.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}
	// Todo: create a verification code
	//session, err := m.UserRepo.DB.StartSession()
	//if err != nil {
	//	fmt.Printf("error while starting session: %v", err)
	//}
	//defer session.EndSession(context.Background())
	//
	//err = mongo.WithSession()
	//
	//err = session.StartTransaction()
	//if err != nil {
	//	fmt.Printf("error creating a transaction: %v", err)
	//	return nil, errors.New("something went wrong")
	//}

	_, err = m.UserRepo.CreateUser(payload)
	if err != nil {
		log.Printf("error while creating user: %v", err)
		return nil, err
	}

	token, err := payload.GenToken()
	if err != nil {
		log.Printf("error while generating password: %v", err)
		return nil, errors.New("something went wrong")
	}
	return &model.AuthResponse{
		AuthToken: token,
		User:      payload,
	}, nil
}

func (m *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := m.UserRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exist")
	}

	err = m.MeetupRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}
	return true, nil
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := m.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	meetup, err = m.MeetupRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}
	return meetup, nil
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}
	meetup := &model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return m.MeetupRepo.CreateMeetup(meetup)
}
