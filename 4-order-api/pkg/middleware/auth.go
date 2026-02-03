package middleware

import (
	"net/http"
	"strings"

	"ps-go-adv/4-order-api/configs"
	"ps-go-adv/4-order-api/pkg/jwt"
	"ps-go-adv/4-order-api/pkg/res"
)

func IsAuthed(next http.Handler,config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if authedHeader == "" || !strings.HasPrefix(authedHeader, "Bearer ") {
				res.Json(w, map[string]string{"error": "unauthorized"}, http.StatusUnauthorized)
				return
			}
			token := strings.TrimPrefix(authedHeader, "Bearer ")
			if token == "" || config.Auth.Secret == "" {
				res.Json(w, map[string]string{"error": "unauthorized"}, http.StatusUnauthorized)
				return
			}
			_, err := jwt.NewJWT(config.Auth.Secret).Validate(token)
			if err != nil {
				res.Json(w, map[string]string{"error": "unauthorized"}, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
	})
}


			