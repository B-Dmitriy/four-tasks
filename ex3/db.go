package main

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const File = "./data.db"

var db *sql.DB

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", File)
	if err != nil {
		return nil, err
	}

	return db, nil
}
