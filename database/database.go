package database

import (
	"database/sql"

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
	return db.Close()
}
