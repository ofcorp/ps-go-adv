package middleware

import (
	"context"
	"net/http"
	"strings"

	"ps-go-adv/4-order-api/configs"
	"ps-go-adv/4-order-api/pkg/jwt"
)

type key string

const (
	ContextPhoneKey key = "ContextEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler,config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if authedHeader == "" || !strings.HasPrefix(authedHeader, "Bearer ") {
				writeUnauthed(w)
				return
			}
			token := strings.TrimPrefix(authedHeader, "Bearer ")
			if token == "" || config.Auth.Secret == "" {
				writeUnauthed(w)
				return
			}
			data, err := jwt.NewJWT(config.Auth.Secret).Validate(token)
			if err != nil {
				writeUnauthed(w)
				return
			}
			ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
	})
}


			