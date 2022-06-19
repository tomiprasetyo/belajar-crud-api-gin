package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal"
)

type createProduct struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Description string  `json:"description" binding:"omitempty,max=250"`
}

// method handler untuk membuat data product
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload createProduct
		var ctx = c.Request.Context()

		// set payload
		if e := c.ShouldBindJSON(&payload); e != nil {
			var res = internal.NewHTTPResponse(http.StatusBadRequest, e)
			c.JSON(http.StatusBadRequest, res)

			return
		}

		// insert payload ke database
		var guid = uuid.New().String()
		var createdAt = time.Now().Format(time.RFC3339)
		if _, e := db.ExecContext(ctx, "INSERT INTO products(guid, name, price, description, createdAt) VALUS(?, ?, ?, ?, ?)", guid, payload.Name, payload.Price, payload.Description, createdAt); e != nil {
			var res = internal.NewHTTPResponse(http.StatusInternalServerError, e)
			c.JSON(http.StatusInternalServerError, res)
			return
		}

		fmt.Println(payload)

		c.String(http.StatusOK, "success create products")
	}
}
