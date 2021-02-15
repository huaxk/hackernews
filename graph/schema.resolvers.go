package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/huaxk/hackernews/graph/generated"
	"github.com/huaxk/hackernews/graph/model"
	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/internal/models"
	"github.com/huaxk/hackernews/pkg/jwt"
)

func (r *linkResolver) UserID(ctx context.Context, obj *models.Link) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*models.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &models.Link{}, fmt.Errorf("access denied")
	}

	link := models.Link{
		Title:   input.Title,
		Address: input.Address,
		UserID:  user.ID,
	}
	r.DB.Create(&link)
	return &link, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{
		Name:     input.Username,
		Password: hashedPassword,
	}
	r.DB.Create(&user)
	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user models.User
	r.DB.Where("name = ?", input.Username).First(&user)
	if correct := models.CheckPasswordHash(input.Password, user.Password); !correct {
		return "", fmt.Errorf("WrongUsernameOrPasswordError")
	}
	token, err := jwt.GenerateToken(user.Name)
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

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	var resultLinks []*models.Link
	r.DB.Preload("User").Find(&resultLinks)

	return resultLinks, nil
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type linkResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
