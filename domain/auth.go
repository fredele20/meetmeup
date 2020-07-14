package domain

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"meetmeup/graph/model"
	"strconv"
	"time"
)

func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := d.UserRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("user with email already exists")
	}

	_, err = d.UserRepo.GetUserByUsername(input.Username)
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

	_, err = d.UserRepo.CreateUser(payload)
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

func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := d.UserRepo.GetUserByEmail(input.Email)
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
