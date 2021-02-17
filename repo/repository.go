package repo

import (
	"context"

	"github.com/huaxk/hackernews/models"
)

type Repository interface {
	CreateLink(ctx context.Context, input models.NewLink) (*models.Link, error)
	CreateUser(ctx context.Context, input models.NewUser) (string, error)
	Login(ctx context.Context, input models.Login) (string, error)
	RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error)
	Links(ctx context.Context) ([]*models.Link, error)
}
