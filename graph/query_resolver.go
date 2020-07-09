package graph

import (
	"context"
	"meetmeup/graph/generated"
	"meetmeup/graph/model"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.GetUserById(id)
}

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	return r.MeetupRepo.GetMeetups(filter, limit, offset)
}
