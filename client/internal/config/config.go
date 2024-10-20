package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	PORT        string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		PORT:        os.Getenv("PORT"),
	}, nil
}
