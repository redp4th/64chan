package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Cookie string `json:"cookie"`
}

func GenerateToken(cookie string, secret []byte) (token string, err error) {
	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "64chan",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		cookie,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(secret)
	return
}

func ParseToken(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(toekn *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
