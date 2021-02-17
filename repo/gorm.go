package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/models"
	"github.com/huaxk/hackernews/pkg/jwt"
	"gorm.io/gorm"
)

type repoSvc struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoSvc{
		db: db,
	}
}

func (r *repoSvc) CreateLink(ctx context.Context, input models.NewLink) (*models.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &models.Link{}, fmt.Errorf("access denied")
	}

	link := models.Link{
		Title:   input.Title,
		Address: input.Address,
		UserID:  user.ID,
	}
	r.db.Create(&link)
	return &link, nil
}

func (r *repoSvc) CreateUser(ctx context.Context, input models.NewUser) (string, error) {
	hashedPassword, err := models.HashPassword(input.Password)
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{
		Name:     input.Username,
		Password: hashedPassword,
	}
	r.db.Create(&user)
	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *repoSvc) Login(ctx context.Context, input models.Login) (string, error) {
	var user models.User
	r.db.Where("name = ?", input.Username).First(&user)
	if correct := models.CheckPasswordHash(input.Password, user.Password); !correct {
		return "", fmt.Errorf("WrongUsernameOrPasswordError")
	}
	token, err := jwt.GenerateToken(user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *repoSvc) RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error) {
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

func (r *repoSvc) Links(ctx context.Context) ([]*models.Link, error) {
	var resultLinks []*models.Link
	r.db.Preload("User").Find(&resultLinks)

	return resultLinks, nil
}
