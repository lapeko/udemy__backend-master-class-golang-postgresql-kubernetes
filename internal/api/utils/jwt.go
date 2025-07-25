package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/lapeko/udemy__backend-master-class-golang-postgresql-kubernetes/internal/api/config"
)

type JWTUserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

var JwtKey = []byte(config.Get().JwtSecretKey)

func ParseJwtToken(tokenString string) (*JWTUserClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTUserClaims{}, func(t *jwt.Token) (any, error) {
		return JwtKey, nil
	})
	if err == nil && token != nil && token.Valid {
		if claims, ok := token.Claims.(*JWTUserClaims); ok {
			return claims, true
		}
	}
	return nil, false
}
