package main

import (
	"4-order-api/configs"
	"4-order-api/internal/auth"
	"4-order-api/internal/product"
	"4-order-api/internal/user"
	"4-order-api/pkg/db"
	"4-order-api/pkg/logger"
	"4-order-api/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	logger.LogInit()
	router := http.NewServeMux()
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	//БД
	userRepository := user.NewUserRepository(db)
	productRepository := product.NewProductRepository(db)

	authService := auth.NewAuthRepository(userRepository)
	//обработчики
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	product.NewProductHandler(router, &product.ProductHandDeps{
		ProductRepository: productRepository,
	})
	server := http.Server{
		Addr:    ":8085",
		Handler: middleware.Logging(router),
	}
	fmt.Println("БД Создана")
	fmt.Printf("Listen port%v\n", server.Addr)
	server.ListenAndServe()

}
