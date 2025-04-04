package config

import (
	"os"
)

type Config struct {
	MasterKey string
}

func NewConfig() *Config {
	MasterKey := os.Getenv("X_MASTER_KEY")
	if MasterKey == "" {
		panic("Пустая строка, нет ключа доступа")
	}

	return &Config{
		MasterKey: MasterKey,
	}
}
