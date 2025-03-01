package repositories

import (
	"database/sql"
	"errors"

	"github.com/lucsky/cuid"
	"github.com/tsraveling/dog-fight/server/internal/db"
)

type Captain struct {
	ID   string // Primary key
	Name string
}

// CaptainRepository defines operations for managing Captain records.
type CaptainRepository interface {
	Create(captain Captain) (string, error)
	Get(id string) (*Captain, error)
	Update(captain Captain) error
	Delete(id string) error
}

// sqliteCaptainRepository is a SQLite‑based implementation of CaptainRepository.
type sqliteCaptainRepository struct {
	db *db.DB
}

var createTableQuery = `
CREATE TABLE IF NOT EXISTS captains (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);`

// NewCaptainRepository creates the captains table (if needed) and returns a repository.
func NewCaptainRepository(db *db.DB) (CaptainRepository, error) {
	if err := db.CreateTable(createTableQuery); err != nil {
		return nil, err
	}
	return &sqliteCaptainRepository{db: db}, nil
}

// Iserts a new captain record.
func (r *sqliteCaptainRepository) Create(captain Captain) (string, error) {
	captain.ID = cuid.New()
	query := `INSERT INTO captains (id, name) VALUES (?, ?);`
	_, err := r.db.Exec(query, captain.ID, captain.Name)
	if err != nil {
		return "", err
	}
	return captain.ID, nil
}

// Get retrieves a captain by ID.
func (r *sqliteCaptainRepository) Get(id string) (*Captain, error) {
	query := `SELECT id, name FROM captains WHERE id = ?;`
	row := r.db.QueryRow(query, id)
	var captain Captain
	err := row.Scan(&captain.ID, &captain.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("captain not found")
	} else if err != nil {
		return nil, err
	}
	return &captain, nil
}

// Update modifies an existing captain record.
func (r *sqliteCaptainRepository) Update(captain Captain) error {
	if !isValidCUID(captain.ID) {
		return errors.New("invalid captain ID: not a valid CUID")
	}

	query := `UPDATE captains SET name = ? WHERE id = ?;`
	result, err := r.db.Exec(query, captain.Name, captain.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("captain not found")
	}
	return nil
}

// Delete removes a captain record by ID.
func (r *sqliteCaptainRepository) Delete(id string) error {
	query := `DELETE FROM captains WHERE id = ?;`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("captain not found")
	}
	return nil
}

func isValidCUID(id string) bool {
	return len(id) == 25 && id[0] == 'c'
}