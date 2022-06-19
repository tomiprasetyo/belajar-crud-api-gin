package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method handler untuk mengupdate data product
func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success update product")
	}
}
