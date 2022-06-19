package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method handler untuk mengambil semua data product
func GetAllProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success get all products")
	}
}
