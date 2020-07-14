package graph

import (
	"context"
	"errors"
	"meetmeup/graph/generated"
	"meetmeup/graph/model"
)

var (
	ErrInput = errors.New("input error")
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.Register(ctx, input)
}

func (m *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.Login(ctx, input)
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return m.Domain.DeleteMeetup(ctx, id)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	return m.Domain.UpdateMeetup(ctx, id, input)
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	return m.Domain.CreateMeetup(ctx, input)
}
