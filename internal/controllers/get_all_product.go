package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal"
)

// method handler untuk mengambil semua data product
func GetAllProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var rows *sql.Rows
		var e error

		// query untuk mengambil field dari table product
		if rows, e = db.Query("SELECT guid, name, price, description, createdAt FROM products"); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		defer rows.Close()
		// slice Product
		var products []Product
		// jika ada rows yg tersedia, set ke Product
		for rows.Next() {
			var product Product

			if e := rows.Scan(&product.GUID, &product.Name, &product.Price, &product.Description, &product.CreatedAt); e != nil {
				var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
				c.JSON(http.StatusInternalServerError, res)
				return
			}

			// jika scan berhasil append product yang beru ke list Product
			products = append(products, product)
		}

		// cek panjang dari product
		if len(products) == 0 {

			var res = internal.NewHTTPResponse(http.StatusNotFound, sql.ErrNoRows)
			c.JSON(http.StatusNotFound, res)
			return
		}

		var res = internal.NewHTTPResponse(http.StatusOK, products)
		c.JSON(http.StatusOK, res)
	}
}
