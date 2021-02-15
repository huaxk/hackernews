package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secret")

type UserClaims struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SecretKey)
}

func ParseToken(tokenStr string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims.Username, nil
		}
	}
	return "", err
}
