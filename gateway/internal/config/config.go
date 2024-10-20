package config

import (
	"os"
)

type Config struct {
	PORT                string
	AuthServiceAddr     string
	ProductsServiceAddr string
	GinMode             string
}

func LoadConfig() (*Config, error) {
	return &Config{
		PORT:                os.Getenv("PORT"),
		AuthServiceAddr:     os.Getenv("auth_service_addr"),
		ProductsServiceAddr: os.Getenv("products_service_addr"),
		GinMode:             os.Getenv("GIN_MODE"),
	}, nil
}
