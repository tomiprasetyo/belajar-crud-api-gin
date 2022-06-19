package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method untuk menghapus data product
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success delete product")
	}
}
