package models

import "golang.org/x/crypto/bcrypt"

type Link struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  string `json:"userID"`
	User    *User  `json:"user" gorm:"foreignKey:UserID"`
}

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
