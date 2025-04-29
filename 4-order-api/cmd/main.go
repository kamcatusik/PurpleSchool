package main

import (
	"4-order-api/configs"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"

	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Run")

	router := http.NewServeMux()
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	productRepository := product.NewProductRepository(db)
	product.NewProductHandler(router, &product.ProductHandDeps{
		ProductRepository: productRepository,
	})
	server := http.Server{
		Addr:    ":8085",
		Handler: router,
	}
	fmt.Println("БД Создана")
	fmt.Printf("Listen port%v\n", server.Addr)
	server.ListenAndServe()

}
