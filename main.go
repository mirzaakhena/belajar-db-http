package main

import (
	"belajar/database"
	"belajar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.OpenDatabase()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// REST API CRUD

	// CREATE
	r.POST("/products", func(c *gin.Context) {

		var product model.Product
		err := c.BindJSON(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = product.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = database.InsertProduct(db, product)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, product)

	})

	// READ ALL
	r.GET("/products", func(c *gin.Context) {
		// ...
	})

	// READ ONE
	r.GET("/products/:product_id", func(c *gin.Context) {
		// ...
	})

	// UPDATE
	r.PUT("/products/:product_id", func(c *gin.Context) {
		// ...
	})

	// DELETE
	r.DELETE("/products/:product_id", func(c *gin.Context) {
		// ...
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
