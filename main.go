package main

import (
	"belajar/database"
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

	fmt.Printf("bisa konek db\n")

	//product := model.Product{
	//	Nama:  "Jeruk",
	//	Harga: 21300,
	//	Stok:  3,
	//}

	//err = database.InsertProduct(db, product)
	//if err != nil {
	//	panic(err)
	//}

	//products, err := database.SelectAllProduct(db)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf("%v\n", products)

	product, err := database.SelectOneProduct(db, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", product)

}
