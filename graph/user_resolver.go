package graph

import (
	"context"
	"meetmeup/graph/generated"
	"meetmeup/graph/model"
)

type userResolver struct{ *Resolver }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

func (u *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	return u.Domain.MeetupRepo.GetMeetupsForUser(obj)
}
