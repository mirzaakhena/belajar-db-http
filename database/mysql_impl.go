package database

import (
	"belajar/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// go get -v github.com/go-sql-driver/mysql

type NativeMySQL struct {
	DB *sql.DB
}

func NewNativeMySQL(db *sql.DB) (*NativeMySQL, error) {

	fmt.Println(">>>>>>>> mysql native implementation")

	return &NativeMySQL{DB: db}, nil
}

func (n NativeMySQL) InsertProduct(product *model.Product) error {

	result, err := n.DB.Exec(`
		INSERT INTO product(nama, harga, stok) 
		VALUES (?, ?, ?)
	`, product.Nama, product.Harga, product.Stok)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = id

	fmt.Printf("product sudah diinsert dengan id %v\n", id)

	return nil
}

func (n NativeMySQL) SelectAllProduct(page, size int) ([]model.Product, error) {

	result := make([]model.Product, 0)

	row, err := n.DB.Query("SELECT * FROM product LIMIT ? OFFSET ?", size, (page-1)*size)
	if err != nil {
		return nil, err
	}

	columns, err := row.Columns()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Columns %v\n", columns)

	for row.Next() {
		err := row.Err()
		if err != nil {
			return nil, err
		}

		var product model.Product

		err = row.Scan(&product.ID, &product.Nama, &product.Harga, &product.Stok)
		if err != nil {
			return nil, err
		}

		result = append(result, product)
	}

	return result, nil
}

func (n NativeMySQL) SelectOneProduct(id int) (*model.Product, error) {

	var product model.Product

	err := n.DB.
		QueryRow("SELECT * FROM product WHERE id = ?", id).
		Scan(&product.ID, &product.Nama, &product.Harga, &product.Stok)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (n NativeMySQL) UpdateOneProduct(product model.Product) error {

	_, err := n.DB.Exec(`
		UPDATE product SET nama = ?, harga = ?, stok = ? WHERE id = ?`,
		product.Nama,
		product.Harga,
		product.Stok,
		product.ID,
	)

	if err != nil {
		return err
	}

	fmt.Printf("product sudah di update\n")

	return nil
}

func (n NativeMySQL) DeleteProduct(id int) error {

	_, err := n.DB.Exec(`DELETE FROM product WHERE id = ?`, id)

	if err != nil {
		return err
	}

	fmt.Printf("product sudah di delete\n")

	return nil
}
