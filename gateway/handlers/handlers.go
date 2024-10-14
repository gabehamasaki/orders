package handlers

import "github.com/gabehamasaki/orders/balancer/config"

type Handler struct {
	Cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		Cfg: cfg,
	}
}
