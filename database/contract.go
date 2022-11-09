package database

import (
	"belajar/model"
	"database/sql"
)

type Crud interface {
	InsertProduct(product *model.Product) error
	SelectAllProduct(page, size int) ([]model.Product, error)
	SelectOneProduct(id int) (*model.Product, error)
	UpdateOneProduct(product model.Product) error
	DeleteProduct(id int) error
}

func OpenDatabase() (*sql.DB, error) {

	db, err := sql.Open(
		"mysql",
		"root:mypass123@tcp(127.0.0.1:3306)/belajar_db",
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDatabase(db *sql.DB) error {
	err := db.Close()
	return err
}
