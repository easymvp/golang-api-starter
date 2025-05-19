package app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtConfig struct {
	Secret         string
	Realm          string
	ExpirationTime int
	RefreshTime    int
}

func DecodeJWTToken(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func EncodeJWTToken(userId string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().AddDate(1, 0, 0).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
