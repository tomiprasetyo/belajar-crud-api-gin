package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk mengambil data product berdasarkan id
func GetProductById(c *gin.Context) {
	c.String(http.StatusOK, "success get product")
}
