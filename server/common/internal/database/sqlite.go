package database

import (
	"database/sql"
	"fmt"
	"time"
)

func openSQLite(path string) (*DB, error) {
	sqliteDB, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite DB: %w", err)
	}

	// Set connection pool settings.
	sqliteDB.SetConnMaxLifetime(5 * time.Minute)
	sqliteDB.SetMaxOpenConns(10)
	sqliteDB.SetMaxIdleConns(5)

	// Verify the connection.
	if err = sqliteDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping SQLite DB: %w", err)
	}

	return &DB{sqliteDB}, nil
}