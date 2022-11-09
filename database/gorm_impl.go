package database

import (
	"belajar/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// go get -v github.com/go-sql-driver/mysql

type GormLib struct {
	DB *gorm.DB
}

func NewGormLiB(db *sql.DB) (*GormLib, error) {

	fmt.Println(">>>>>>>> gorm implementation")

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = gormDB.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, err
	}

	return &GormLib{DB: gormDB}, nil
}

func (n GormLib) InsertProduct(product *model.Product) error {

	err := n.DB.Save(product).Error
	if err != nil {
		return err
	}

	fmt.Printf("product sudah di insert\n")

	return nil
}

func (n GormLib) SelectAllProduct(page, size int) ([]model.Product, error) {

	products := make([]model.Product, 0)

	err := n.DB.Limit(size).Offset((page - 1) * size).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (n GormLib) SelectOneProduct(id int) (*model.Product, error) {

	var product model.Product

	err := n.DB.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (n GormLib) UpdateOneProduct(product model.Product) error {

	err := n.DB.Save(product).Error
	if err != nil {
		return err
	}

	fmt.Printf("product sudah di update\n")

	return nil
}

func (n GormLib) DeleteProduct(id int) error {

	err := n.DB.Unscoped().Delete(&model.Product{}, id).Error
	if err != nil {
		return err
	}

	fmt.Printf("product sudah di delete\n")

	return nil
}
