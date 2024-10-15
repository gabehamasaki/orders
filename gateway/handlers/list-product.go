package handlers

import (
	"net/http"
	"strconv"
	"strings"

	pb "github.com/gabehamasaki/orders/grpc/pb/proto/v1"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListProducts(c *gin.Context) {

	page := c.Query("page")
	perPage := c.Query("per_page")

	// Set default values
	pageInt := int64(1)     // Default page is 1
	perPageInt := int64(15) // Default per_page is 15

	// Parse page if it exists
	if page != "" {
		var err error
		pageInt, err = strconv.ParseInt(page, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}
	}

	// Parse per_page if it exists
	if perPage != "" {
		var err error
		perPageInt, err = strconv.ParseInt(perPage, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid per_page"})
			return
		}
	}

	res, err := h.client.ListProducts(c, &pb.ListProductsRequest{
		Page:    int32(pageInt),
		PerPage: int32(perPageInt),
	})

	if err != nil {
		errArray := strings.Split(err.Error(), "desc = ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errArray[len(errArray)-1],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":          res.Products,
		"total":         res.Total,
		"total_pages":   res.TotalPages,
		"page":          pageInt,
		"total_in_page": len(res.Products),
	})
}
