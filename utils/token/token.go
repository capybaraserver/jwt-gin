package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

var secretKey = []byte("HELLOWORLD")

func GenerateJWT(userID int64) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func keyFunc(_ *jwt.Token) (interface{}, error) {
	return secretKey, nil
}

func ParseJWT(tokenString string) (*MyClaims, error) {
	var claims MyClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims, nil
}
