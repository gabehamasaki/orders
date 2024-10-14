package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT            string
	AuthServiceAddr string
	GinMode         string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		PORT:            os.Getenv("PORT"),
		AuthServiceAddr: os.Getenv("auth_service_addr"),
		GinMode:         os.Getenv("GIN_MODE"),
	}, nil
}
