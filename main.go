package main

import (
	"belajar/database"
	"belajar/server"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.OpenDatabase()
	if err != nil {
		panic(err)
	}

	service := server.MyService{
		DB: db,
	}

	r := gin.Default()

	// REST API CRUD

	// CREATE
	r.POST("/products", service.CreateProduct)

	// READ ALL
	r.GET("/products", service.GetAllProduct)

	// READ ONE
	r.GET("/products/:product_id", service.GetOneProduct)

	// UPDATE
	r.PUT("/products/:product_id", service.UpdateProduct)

	// DELETE
	r.DELETE("/products/:product_id", service.DeleteProduct)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
