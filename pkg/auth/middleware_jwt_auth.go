package auth

import (
	"context"
	"github.com/djumanoff/gotodo/pkg/http-helper"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type MiddlewareFactory interface {
	JWTParser(handler http_helper.Handler) http_helper.Handler
}

type auth struct {
	pubKey string
	errSys http_helper.ErrorSystem
}

func NewAuthMW(pubKey string, errSys http_helper.ErrorSystem) MiddlewareFactory {
	return &auth{pubKey: pubKey, errSys: errSys}
}

func (auth *auth) JWTParser(handler http_helper.Handler) http_helper.Handler {
	return func(ctx context.Context, r *http.Request) http_helper.Response {
		header := r.Header.Get("Authorization")
		if header == "" {
			return auth.errSys.Unauthorized(10, "Unauthorized.")
		}
		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return auth.errSys.Unauthorized(11, "Unauthorized.")
		}
		jwtToken := parts[1]
		if jwtToken == "" {
			return auth.errSys.Unauthorized(12, "Unauthorized.")
		}
		claims := &JWTClaims{}
		_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.pubKey), nil
		})
		if err != nil {
			return auth.errSys.Forbidden(13, err.Error())
		}
		ctx = context.WithValue(ctx, "claims", claims)
		return handler(ctx, r)
	}
}
