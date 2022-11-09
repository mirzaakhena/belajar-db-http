package server

import (
	"belajar/database"
	"belajar/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MyService struct {
	Repo database.Crud
}

func (m MyService) CreateProduct(c *gin.Context) {

	var product model.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = product.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = m.Repo.InsertProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: product, ErrorMessage: ""})

}

func (m MyService) GetAllProduct(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	sizeStr := c.DefaultQuery("size", "3")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	products, err := m.Repo.SelectAllProduct(page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: products, ErrorMessage: ""})

}

func (m MyService) GetOneProduct(c *gin.Context) {

	idStr := c.Param("product_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	product, err := m.Repo.SelectOneProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: product, ErrorMessage: ""})

}

func (m MyService) UpdateProduct(c *gin.Context) {

	idStr := c.Param("product_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	var product model.Product
	err = c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = product.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	product.ID = int64(id)

	err = m.Repo.UpdateOneProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: nil, ErrorMessage: ""})

}

func (m MyService) DeleteProduct(c *gin.Context) {

	idStr := c.Param("product_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = m.Repo.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: nil, ErrorMessage: ""})

}

type Response struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"error_message"`
}
