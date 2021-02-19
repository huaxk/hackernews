package models

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

var DbModels = []interface{}{
	&User{},
	&Link{},
}

type Link struct {
	ID      uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title   string    `json:"title"`
	Address string    `json:"address"`
	UserID  uuid.UUID `json:"userID"`
	User    User      `json:"user"`
}

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
