package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk mengupdate data product
func UpdateProduct(c *gin.Context) {
	c.String(http.StatusOK, "success update product")
}
