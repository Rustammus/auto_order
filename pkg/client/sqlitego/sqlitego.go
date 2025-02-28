package sqlitego

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
