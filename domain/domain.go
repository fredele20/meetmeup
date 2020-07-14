package domain

import (
	"errors"
	"meetmeup/database"
	"meetmeup/graph/model"
)

var (
	ErrBadCredentials  = errors.New("email/password combination does not match")
	ErrUnAuthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("forbidden, unauthorized  ")
)

type Domain struct {
	UserRepo   database.UsersRepo
	MeetupRepo database.MeetupsRepo
}

func NewDomain(userRepo database.UsersRepo, meetupRepo database.MeetupsRepo) *Domain {
	return &Domain{UserRepo: userRepo, MeetupRepo: meetupRepo}
}

type Ownable interface {
	IsOwner(user *model.User) bool
}

func checkOwnerShip(o Ownable, user *model.User) bool {
	return o.IsOwner(user)
}
