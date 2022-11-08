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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = database.CloseDatabase(db)
	if err != nil {
		panic(err)
	}

	fmt.Printf("bisa konek db\n")

}
