package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk menghapus data product
func DeleteProduct(c *gin.Context) {
	c.String(http.StatusOK, "success delete products")
}
