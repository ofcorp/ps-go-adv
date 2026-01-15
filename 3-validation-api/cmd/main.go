package main

import (
	"fmt"
	"net/http"
	"ps-go-adv/3-validation-api/internal/verify"
)

func main() {
	// conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
