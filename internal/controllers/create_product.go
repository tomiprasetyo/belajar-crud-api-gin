package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk membuat data product
func CreateProduct(c *gin.Context) {
	c.String(http.StatusOK, "success create product")
}
