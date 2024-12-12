package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"t0/model"
	"time"
)

var jwtKey = []byte("love_library")

func GenJwt(name string) string {
	jc := model.JwtClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "_jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)

	ss, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}
	return ss
}
