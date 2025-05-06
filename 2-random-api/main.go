package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewRandHandler(router)
	server := http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	fmt.Println("Server run")
	server.ListenAndServe()
}
