package middleware

import (
	"net/http"
	"strings"

	"github.com/gabehamasaki/orders/gateway/clients"
	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	client *clients.Client
}

func NewMiddleware(client *clients.Client) *Middleware {
	return &Middleware{client: client}
}

func (m *Middleware) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authorizationHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authorizationHeader) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		token := authorizationHeader[1]

		res, err := m.client.ValidateToken(c, &pb.ValidateTokenRequest{Token: token})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !res.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID in the context
		c.Set("userID", res.Id)

		// Continue to the next middleware
		c.Next()
	}
}
