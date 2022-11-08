package database

import (
	"belajar/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go get -v github.com/go-sql-driver/mysql

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

func InsertProduct(db *sql.DB, product model.Product) error {

	row, err := db.Query(`
		INSERT INTO product(nama, harga, stok) 
		VALUES (?, ?, ?)
	`, product.Nama, product.Harga, product.Stok)

	if err != nil {
		return err
	}

	err = row.Close()
	if err != nil {
		return err
	}

	fmt.Printf("product sudah diinsert\n")

	return nil

}
