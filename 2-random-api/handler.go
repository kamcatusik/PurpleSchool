package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type RandHandler struct{}

func NewRandHandler(router *http.ServeMux) {
	handler := &RandHandler{}
	router.HandleFunc("/rand", handler.randNum)
}
func (handler *RandHandler) randNum(w http.ResponseWriter, req *http.Request) {
	random := strconv.Itoa(rand.Intn(7) + 1)
	_, err := w.Write([]byte(random))
	if err != nil {
		fmt.Println("не удалось записать ответ")
	}

}
