package main

import (
	"4-order-api/cmd/pkg/db"
	"4-order-api/configs"
	"fmt"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	fmt.Println("БД Создана")
}
