package handlers

import (
	"github.com/gabehamasaki/orders/gateway/clients"
	"github.com/gabehamasaki/orders/gateway/config"
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
