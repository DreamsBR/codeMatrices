package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("mySecretKey")

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
