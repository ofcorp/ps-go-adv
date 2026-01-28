package main

import (
	"fmt"
	"ps-go-adv/4-order-api/configs"
	"ps-go-adv/4-order-api/internal/product"
	"ps-go-adv/4-order-api/pkg/db"
	"ps-go-adv/4-order-api/pkg/middleware"

	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	productRepository := product.NewProductRepository(db)
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
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