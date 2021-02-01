package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/huaxk/hackernews/graph/generated"
	"github.com/huaxk/hackernews/graph/model"
	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/internal/links"
	"github.com/huaxk/hackernews/internal/users"
	"github.com/huaxk/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	link := links.Link{
		Title:   input.Title,
		Address: input.Address,
		User:    user,
	}
	linkID := link.Save()
	return &model.Link{
		ID:      strconv.FormatInt(linkID, 10),
		Title:   link.Title,
		Address: link.Address,
		User: &model.User{
			ID:   user.ID,
			Name: user.Username,
		},
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	if correct := user.Authenticate(); !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied: %s", err)
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	dbLinks := links.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
			User: &model.User{
				ID:   link.User.ID,
				Name: link.User.Username,
			},
		})
	}
	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
