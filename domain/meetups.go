package domain

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"meetmeup/graph/model"
	"meetmeup/middleware"
	"strconv"
)

func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, ErrUnAuthenticated
	}
	meetup, err := d.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exist")
	}

	if !meetup.IsOwner(currentUser) {
		return false, ErrForbidden
	}

	err = d.MeetupRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}
	return true, nil
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnAuthenticated
	}

	meetup, err := d.MeetupRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	if !meetup.IsOwner(currentUser) {
		return nil, ErrForbidden
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

	meetup, err = d.MeetupRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}
	return meetup, nil
}

func (d *Domain) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
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
		ID:          strconv.Itoa(rand.Int()),
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return d.MeetupRepo.CreateMeetup(meetup)
}
