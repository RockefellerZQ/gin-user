package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("gin_user_secret")

type TokenClaims struct {
	UserId uint
	jwt.StandardClaims
}

func GenerateToken(userId uint) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := TokenClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:   time.Now().Unix(),
			Issuer:    "com.ant",
			Subject:   "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseTime(tokenString string) (*jwt.Token, *TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
