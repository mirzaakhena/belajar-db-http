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

	// stmt, err := db.Prepare("INSERT INTO product(nama, harga, stok) VALUES (?, ?, ?)")
	// if err != nil {
	// 	return err
	// }
	// stmt.Exec(product.Nama, product.Harga, product.Stok)

	result, err := db.Exec(`
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

	fmt.Printf("product sudah diinsert dengan id %v\n", id)

	return nil

}

func SelectAllProduct(db *sql.DB) ([]model.Product, error) {

	result := make([]model.Product, 0)

	row, err := db.Query("SELECT * FROM product")
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

func SelectOneProduct(db *sql.DB, id int) (*model.Product, error) {

	var product model.Product

	err := db.
		QueryRow("SELECT * FROM product WHERE id = ?", id).
		Scan(&product.ID, &product.Nama, &product.Harga, &product.Stok)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func UpdateOneProduct(db *sql.DB, product model.Product) error {

	_, err := db.Exec(`
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

func DeleteProduct(db *sql.DB, id int) error {

	_, err := db.Exec(`DELETE product WHERE id = ?`, id)

	if err != nil {
		return err
	}

	fmt.Printf("product sudah di delete\n")

	return nil
}
