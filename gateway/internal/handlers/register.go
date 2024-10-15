package handlers

import (
	"net/http"
	"strings"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(c *gin.Context) {
	body := &RegisterRequest{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	_, err := h.client.Register(c, &pb.RegisterRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password, // Fixed to use the correct password field
	})
	if err != nil {
		errArray := strings.Split(err.Error(), "desc = ")
		switch errArray[len(errArray)-1] {
		case "User already exists":
			c.JSON(http.StatusConflict, gin.H{
				"error": errArray[len(errArray)-1],
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errArray[len(errArray)-1],
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered",
	})
}
