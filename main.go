package main

import (
	"belajar/database"
	"belajar/model"
	"fmt"
)

func main() {

	db, err := database.OpenDatabase()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = database.CloseDatabase(db)
		if err != nil {
			panic(err)
		}
	}()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	product := model.Product{
		Nama:  "Manggis",
		Harga: 13600,
		Stok:  9,
	}

	err = database.InsertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Printf("bisa konek db\n")

}
