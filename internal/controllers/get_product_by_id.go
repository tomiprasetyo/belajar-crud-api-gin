package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method handler untuk mengambil data product berdasarkan id
func GetProductById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success get product ")
	}
}
