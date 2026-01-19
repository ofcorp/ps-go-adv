package verify

import (
	"net/http"
	"net/smtp"
	"ps-go-adv/3-validation-api/configs"
	"ps-go-adv/3-validation-api/pkg/common"
	"ps-go-adv/3-validation-api/pkg/req"
	"ps-go-adv/3-validation-api/pkg/res"
	"ps-go-adv/3-validation-api/repository"
	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct{
	*configs.Config
	*repository.Storage
}

type VerifyHandler struct{
	*configs.Config
	*repository.Storage
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
		Storage: deps.Storage,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}
		hash, err := common.UniqueHash()
		if err != nil {
			return
		}
		err = handler.Storage.AddItem(body.Email, hash)
		if err != nil {
			return
		}
		e := email.NewEmail()
		e.From = handler.Mailer.Email
		e.To = []string{body.Email}
		e.Subject = "Verification Email"
		e.HTML = []byte("Click <a href='http://localhost:8081/verify/" + hash + "'>here</a> to verify your email.")
		e.Send(handler.Mailer.Host, smtp.PlainAuth("", handler.Mailer.Email, handler.Mailer.Password, handler.Mailer.Host))
		data := SendResponse{
			Hash: hash,
		}
		res.Json(w, data, 201)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		result := handler.Storage.VerifyHash(hash)
		data := VerifyResponse{
			Result: result,
		}
		res.Json(w, data, 200)		
	}
}