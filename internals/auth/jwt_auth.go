package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(name string) string {
	tokenSecret := os.Getenv("TOKEN_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{})
	return ""
}
