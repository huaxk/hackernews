package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/huaxk/hackernews/graph/generated"
	"github.com/huaxk/hackernews/models"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input models.NewLink) (*models.Link, error) {
	return r.Repository.CreateLink(ctx, input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (string, error) {
	return r.Repository.CreateUser(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (string, error) {
	return r.Repository.Login(ctx, input)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error) {
	return r.Repository.RefreshToken(ctx, input)
}

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	return r.Repository.Links(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
