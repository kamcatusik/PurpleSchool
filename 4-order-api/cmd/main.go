package main

import (
	"4-order-api/configs"
	"4-order-api/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	server := http.Server{
		Addr:    ":8085",
		Handler: router,
	}
	fmt.Println("БД Создана")
	fmt.Printf("Listen port%v\n", server.Addr)
	server.ListenAndServe()

}
