package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

// OpenDB opens (or creates) the SQLite database at the specified path.
func OpenDB(path string) (*DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Verify connection.
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// CreateTable executes a table creation query.
func (db *DB) CreateTable(query string) error {
	_, err := db.Exec(query)
	return err
}

