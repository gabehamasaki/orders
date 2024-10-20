package config

import (
	"os"
)

type Config struct {
	DatabaseURL  string
	PORT         string
	JWTSecretKey string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		PORT:         os.Getenv("PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}, nil
}
