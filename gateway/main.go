package main

import (
	"fmt"
	"net/http"

	"github.com/gabehamasaki/orders/gateway/clients"
	"github.com/gabehamasaki/orders/gateway/config"
	"github.com/gabehamasaki/orders/gateway/handlers"
	"github.com/gabehamasaki/orders/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	gin.SetMode(cfg.GinMode)

	logger, err := logger.NewLogger("gateway")
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	r := gin.Default()

	handlers := handlers.NewHandler(cfg, clients.NewClient(logger, cfg))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", handlers.Register)

	logger.Info("Server starting", zap.String("address", fmt.Sprintf("0.0.0.0:%s", cfg.PORT)))
	r.Run(fmt.Sprintf("0.0.0.0:%s", cfg.PORT))
}
