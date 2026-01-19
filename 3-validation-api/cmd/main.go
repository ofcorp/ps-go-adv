package main

import (
	"fmt"
	"net/http"
	"ps-go-adv/3-validation-api/configs"
	"ps-go-adv/3-validation-api/internal/verify"
	"ps-go-adv/3-validation-api/repository"
)

func main() {
	conf := configs.LoadConfig()
	vault := repository.NewStorage()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		Config:  conf,
		Storage: vault,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
