package handlers

import (
	"github.com/gabehamasaki/orders/gateway/internal/clients"
	"github.com/gabehamasaki/orders/gateway/internal/config"
)

type Handler struct {
	cfg    *config.Config
	client *clients.Client
}

func NewHandler(cfg *config.Config, client *clients.Client) *Handler {
	return &Handler{
		cfg:    cfg,
		client: client,
	}
}
