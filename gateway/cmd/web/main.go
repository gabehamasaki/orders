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

	api := r.Group("/api")
	{
		api.POST("/auth/login", handlers.Login)
		private := r.Group("/api", middleware.Authenticated())
		{
			private.POST("/products", handlers.CreateProduct)
			private.GET("/products", handlers.ListProducts)
			private.GET("/products/:id", handlers.GetProduct)
		}
	}

	routes := r.Routes()
	for _, route := range routes {
		logger.Info("Registered route", zap.String("method", route.Method), zap.String("path", route.Path))
	}

	logger.Info("Starting HTTP server", zap.String("address", fmt.Sprintf(":%s", cfg.PORT)))
	r.Run(fmt.Sprintf(":%s", cfg.PORT))
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
