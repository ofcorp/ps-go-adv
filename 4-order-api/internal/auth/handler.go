package auth

import (
	"net/http"
	"ps-go-adv/4-order-api/configs"
	"ps-go-adv/4-order-api/pkg/jwt"
	"ps-go-adv/4-order-api/pkg/req"
	"ps-go-adv/4-order-api/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/verify", handler.Verify())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		sessionID, err := handler.AuthService.Login(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			SessionID: sessionID,
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			return
		}
		phone, err := handler.AuthService.Verify(body.SessionID, body.VerificationCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := VerifyResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}
