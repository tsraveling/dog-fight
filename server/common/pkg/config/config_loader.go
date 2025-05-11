package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig holds database configuration details.
type DBConfig struct {
	Driver 				string
	PostgresConnection 	string
	SQLiteFilePath 		string
}

// LoadEnv loads environment variables and returns a DBConfig.
func LoadEnv() (*DBConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		return nil, fmt.Errorf("DB_DRIVER is not set in .env")
	}

	config := &DBConfig{Driver: driver}

	switch driver {
	case "postgres":
		conn := os.Getenv("POSTGRES_CONNECTION_STRING")
		if conn == "" {
			return nil, fmt.Errorf("POSTGRES_CONNECTION_STRING is not set in .env")
		}
		config.PostgresConnection = conn
	case "sqlite":
		path := os.Getenv("SQLITE_FILE_PATH")
		if path == "" {
			return nil, fmt.Errorf("SQLITE_FILE_PATH is not set in .env")
		}
		config.SQLiteFilePath = path
	default:
		return nil, fmt.Errorf("unsupported DB_DRIVER: %s", driver)
	}

	return config, nil
}