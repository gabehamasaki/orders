package handlers

import (
	"net/http"
	"strings"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct { // Fixed the struct definition
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(c *gin.Context) {
	body := &LoginRequest{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	res, err := h.client.Login(c, &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		errArray := strings.Split(err.Error(), "desc = ")
		switch errArray[len(errArray)-1] {
		case "User not found":
			c.JSON(http.StatusNotFound, gin.H{
				"error": errArray[len(errArray)-1],
			})
		case "Invalid password":
			c.JSON(http.StatusUnauthorized, gin.H{
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
		"token": res.Token,
	})
}
