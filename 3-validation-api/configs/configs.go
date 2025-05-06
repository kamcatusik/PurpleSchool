package configs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
}

func LoadConfig() *Config {
	fmt.Println(os.Getwd())
	data, err := os.ReadFile("email.json")
	if err != nil {
		log.Println("Не удалось прочитать файл с email")
	}
	var emailConfig Config
	json.Unmarshal(data, &emailConfig)
	return &emailConfig

}
