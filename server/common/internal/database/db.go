package database

import (
	"database/sql"
	"fmt"

	"github.com/tsraveling/dog-fight/server/common/pkg/config"
)

// DB wraps an sql.DB connection.
type DB struct {
	*sql.DB
}

// InitDB initializes the database based on the configuration.
func InitDB(cfg *config.DBConfig) (*DB, error) {
	switch cfg.Driver {
	case "postgres":
		return openPostgres(cfg.PostgresConnection)
	case "sqlite":
		return openSQLite(cfg.SQLiteFilePath)
	default:
		return nil, fmt.Errorf("unsupported DB_DRIVER: %s", cfg.Driver)
	}
}

// CreateTable executes the provided SQL query to create a table.
func (db *DB) CreateTable(query string) error {
	_, err := db.Exec(query)
	return err
}