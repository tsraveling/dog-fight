package database

import (
	"database/sql"
	"fmt"
	"time"
)

func openPostgres(connStr string) (*DB, error) {
	pgDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL DB: %w", err)
	}

	// Set connection pool settings.
	pgDB.SetConnMaxLifetime(5 * time.Minute)
	pgDB.SetMaxOpenConns(10)
	pgDB.SetMaxIdleConns(5)

	// Verify the connection.
	if err = pgDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL DB: %w", err)
	}

	return &DB{pgDB}, nil
}