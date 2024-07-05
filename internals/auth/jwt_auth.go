package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(name string) string {
	tokenSecret := os.Getenv("TOKEN_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject: name,
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		fmt.Println("couldnt create token")
		return ""
	}

	return tokenString
}

func ValidateToken(tokenString, tokenSecret string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		fmt.Println("couldnt parse the token")
		return
	}
	if !token.Valid {
		return
	}
}
