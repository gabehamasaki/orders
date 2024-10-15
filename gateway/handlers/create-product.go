package handlers

import (
	"net/http"
	"strings"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageUrl    string  `json:"image_url"`
}

func (h *Handler) CreateProduct(c *gin.Context) {
	body := &CreateProductRequest{}
	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	_, err := h.client.CreateProduct(c, &pb.CreateProductRequest{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		ImageUrl:    body.ImageUrl,
	})
	if err != nil {
		errArray := strings.Split(err.Error(), "desc = ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errArray[len(errArray)-1],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product created",
	})
}
