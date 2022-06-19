package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal/controllers"
)

func main() {

	var router = gin.Default()

	// alamat di localhost
	var address = ":3000"

	var db *sql.DB
	var e error

	// koneksi ke database
	if db, e = sql.Open("sqlite3", "/data.db"); e != nil {
		log.Fatalf("Error : %v", e)
	}
	defer db.Close()

	if e := db.Ping(); e != nil {
		log.Fatalf("Error : %v", e)
	}

	// endpoint api
	router.GET("/products", controllers.GetAllProduct(db))
	router.GET("/product/:guid", controllers.GetProductById(db))
	router.POST("/product", controllers.CreateProduct(db))
	router.PUT("/product/:guid", controllers.UpdateProduct(db))
	router.DELETE("/product/:guid", controllers.DeleteProduct(db))

	log.Fatalln(router.Run(address))
}
