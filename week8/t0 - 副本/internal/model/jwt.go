package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Name             string
	RegisteredClaims jwt.RegisteredClaims
}

func (jc JwtClaims) Valid() error {
	return nil
}
