package graph

import (
	"context"
	"meetmeup/graph/generated"
	"meetmeup/graph/model"
)

type meetupResolver struct{ *Resolver }

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

func (m *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}
