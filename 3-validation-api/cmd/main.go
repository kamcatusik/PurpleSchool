package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
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
	fmt.Println("Работа завершена")
}
