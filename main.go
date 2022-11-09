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

	defer func() {
		err := database.CloseDatabase(db)
		if err != nil {
			panic(err)
		}
	}()

	service := server.MyService{DB: db}

	r := gin.Default()

	r.POST("/products", service.CreateProduct)
	r.GET("/products", service.GetAllProduct)
	r.GET("/products/:product_id", service.GetOneProduct)
	r.PUT("/products/:product_id", service.UpdateProduct)
	r.DELETE("/products/:product_id", service.DeleteProduct)

	r.Run()

}
