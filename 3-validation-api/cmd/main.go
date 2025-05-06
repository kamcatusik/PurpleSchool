package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/verify"
	"3-validation-api/logger"
	"fmt"
	"net/http"
)

func main() {
	logger.LogInit()
	fmt.Println("Запуск")
	email := configs.LoadConfig()
	fmt.Println(email.Address, email.Email, email.Password)
	router := http.NewServeMux()
	verify.NewEmailHandler(router, verify.EmailHandler{
		Config: email,
	})
	server := http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	fmt.Printf("Listen port%v\n", server.Addr)
	server.ListenAndServe()

}
