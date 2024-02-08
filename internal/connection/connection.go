package connection

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}

	return db
}

func Init() error {
	db, err := sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS redirects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			path VARCHAR NOT NULL,
			url VARCHAR NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		return err
	}

	return nil
}
