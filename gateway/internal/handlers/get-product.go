package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProduct(c *gin.Context) {

	productID := c.Param("id")

	if productID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	res, err := h.client.GetProduct(c, productID)
	if err != nil {
		errArray := strings.Split(err.Error(), "desc = ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errArray[len(errArray)-1],
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
