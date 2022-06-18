package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tomiprasetyo/belajar-crud-api-gin/internal/controllers"
)

func main() {

	var router = gin.Default()
	var address = ":3000"

	// endpoint api
	router.GET("/products", controllers.GetAllProduct)
	router.GET("/product/:id", controllers.GetProductById)
	router.POST("/product/", controllers.CreateProduct)
	router.PUT("/product/:id", controllers.UpdateProduct)
	router.DELETE("/product/:id", controllers.DeleteProduct)

	log.Fatalln(router.Run(address))
}
