package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func GetSQLiteDB(path string) (*sql.DB, error) {
	if _, err := os.Stat(path); err != nil {
		_, err = os.Create(path)

		if err != nil {
			return nil, err
		}
	}
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	return db, nil
}
