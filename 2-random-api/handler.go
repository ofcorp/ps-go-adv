package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"net/http"
)

type RandHandler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &RandHandler{}
	router.HandleFunc("/rand", handler.Hello())
}

func (handler *RandHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		x, err := crand.Int(crand.Reader, big.NewInt(6))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		n := int(x.Int64()) + 1
		w.Write([]byte(fmt.Sprintf("%d", n)))
	}
}
