package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// method handler untuk membuat data product
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success create products")
	}
}
