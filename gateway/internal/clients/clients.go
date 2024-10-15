package clients

import (
	"github.com/gabehamasaki/orders/gateway/internal/config"
	"go.uber.org/zap"
)

type Client struct {
	logger *zap.Logger
	cfg    *config.Config
}

func NewClient(logger *zap.Logger, cfg *config.Config) *Client {
	return &Client{
		logger: logger,
		cfg:    cfg,
	}
}
