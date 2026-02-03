package main

import (
	"fmt"
	"ps-go-adv/4-order-api/configs"
	"ps-go-adv/4-order-api/internal/auth"
	"ps-go-adv/4-order-api/internal/product"
	"ps-go-adv/4-order-api/internal/user"
	"ps-go-adv/4-order-api/pkg/db"
	"ps-go-adv/4-order-api/pkg/middleware"

	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	productRepository := product.NewProductRepository(db)
	userRepository := user.NewUserRepository(db)

	authService := auth.NewAuthService(userRepository)

	product.NewProductHandler(router, product.ProductHandlerDeps{
		Config:         conf,
		ProductRepository: productRepository,
	})

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}