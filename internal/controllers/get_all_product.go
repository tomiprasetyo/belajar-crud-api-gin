package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk mengambil semua data product
func GetAllProduct(c *gin.Context) {
	c.String(http.StatusOK, "success get all products")
}
