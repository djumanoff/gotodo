package auth

import (
	"github.com/djumanoff/gotodo/pkg/http-helper"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type MiddlewareFactory interface {
	JWT(handler http_helper.Handler) http_helper.Handler
}

type auth struct {
	pubKey string
	errSys http_helper.ErrorSystem
}

func NewAuthMW(pubKey string, errSys http_helper.ErrorSystem) AuthMiddlewareFactory {
	return &auth{pubKey: pubKey, errSys: errSys}
}

func (auth *auth) JWT(handler http_helper.Handler) http_helper.Handler {
	return func(r *http.Request) http_helper.Response {
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
		tkn, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return auth.pubKey, nil
		})
		if err != nil {
			return auth.errSys.Unauthorized(13, err.Error())
		}

		return handler(r)
	}
}
