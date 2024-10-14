package main

import (
	"fmt"
	"net/http"

	"github.com/gabehamasaki/orders/balancer/config"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	gin.SetMode(cfg.GinMode)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf("0.0.0.0:%s", cfg.PORT))
}
