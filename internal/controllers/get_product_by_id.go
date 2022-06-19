package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal"
)

type guidBinding struct {
	GUID string `uri:"guid" binding:"required,uuid4"`
}

// method handler untuk mengambil data product berdasarkan id
func GetProductById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass guid dari uri agar bisa dicari di database
		var binding guidBinding
		var ctx = c.Request.Context()

		if e := c.ShouldBindUri(&binding); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		var row = db.QueryRowContext(ctx, "SELECT guid, name, price, description, createdAt FROM products WHERE guid=?", binding.GUID)

		var product Product
		if e := row.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
			if e == sql.ErrNoRows {
				var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
				c.JSON(http.StatusNotFound, res)
				return
			}

			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		var res = internal.NewHTTPResponse(http.StatusOK, product)
		c.JSON(http.StatusOK, res)
	}
}
