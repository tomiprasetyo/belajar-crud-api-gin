package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal"
)

// method untuk menghapus data product
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// pass guid dari uri agar bisa dicari di database
		var binding guidBinding
		var ctx = c.Request.Context()

		if e := c.ShouldBindUri(&binding); e != nil {
			var res = internal.NewHTTPResponse(http.StatusBadRequest, e)
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// query untuk menghapus data product
		var result sql.Result
		var e error
		if result, e = db.ExecContext(ctx, "DELETE FROM products WHERE guid=?", binding.GUID); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		if nProducts, _ := result.RowsAffected(); nProducts == 0 {
			var res = internal.NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(http.StatusNotFound, res)
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
