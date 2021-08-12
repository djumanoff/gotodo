package auth

import (
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims
}

func (c JWTClaims) Valid() error {
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}
	return nil
}
