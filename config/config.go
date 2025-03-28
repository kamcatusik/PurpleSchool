package config

import (
	"cli/jason/logger"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		logger.InfoLog.Fatalf("Не передан KEY")
	}
	return &Config{
		Key: key,
	}
}
