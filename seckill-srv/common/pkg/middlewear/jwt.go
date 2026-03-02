package middlewear

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "github.com/golang-jwt/jwt/v5"

func GenerateAccessToken(userId int64) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"userId": userId,
		"iat":    now.Unix(),
		"exp":    now.Add(24 * time.Hour).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	switch {
	case token.Valid:
		fmt.Println("You look nice today")
	case errors.Is(err, jwt.ErrTokenMalformed):
		fmt.Println("That's not even a token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		fmt.Println("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	default:
		fmt.Println("Couldn't handle this token:", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, err
	} else {
		fmt.Println(err)
	}
	return nil, nil
}
