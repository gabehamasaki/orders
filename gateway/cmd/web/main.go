package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gabehamasaki/orders/gateway/internal/clients"
	"github.com/gabehamasaki/orders/gateway/internal/config"
	"github.com/gabehamasaki/orders/gateway/internal/handlers"
	"github.com/gabehamasaki/orders/gateway/internal/middleware"
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

	client := clients.NewClient(logger, cfg)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(LoggerMiddleware(logger))

	handlers := handlers.NewHandler(cfg, client)
	middleware := middleware.NewMiddleware(client)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)
	r.POST("/products", middleware.Authenticated(), handlers.CreateProduct)
	r.GET("/products", middleware.Authenticated(), handlers.ListProducts)

	logger.Info("Starting HTTP server", zap.String("address", fmt.Sprintf("0.0.0.0:%s", cfg.PORT)), zap.String("mode", cfg.GinMode))
	r.Run(fmt.Sprintf("0.0.0.0:%s", cfg.PORT))
}

// LoggerMiddleware logs the request and response of the API
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log request received
		logger.Info("Request Received", zap.String("method", c.Request.Method), zap.String("path", c.Request.URL.Path))

		// Process request
		c.Next()

		// Log request completion
		logger.Info("Request Completed",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
		)
	}

}
