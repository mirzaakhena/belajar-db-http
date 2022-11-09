package server

import (
	"belajar/database"
	"belajar/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyService struct {
	DB *sql.DB
}

func (m MyService) CreateProduct(c *gin.Context) {

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

	err = database.InsertProduct(m.DB, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)

}

func (m MyService) GetAllProduct(c *gin.Context) {
	// ...
}

func (m MyService) GetOneProduct(c *gin.Context) {
	// ...
}

func (m MyService) UpdateProduct(c *gin.Context) {
	// ...
}

func (m MyService) DeleteProduct(c *gin.Context) {
	// ...
}
