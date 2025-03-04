package repositories

import (
	"database/sql"
	"errors"

	"github.com/lucsky/cuid"
	"github.com/tsraveling/dog-fight/server/internal/db"
)

// Define the Captain struct.
type Captain struct {
	ID   			string 	// Primary key
	Username 		string 	// Unique username for this server
	PasswordHash 	string	// Hashed password for authentication
	Rank			string 	// Captain's rank (e.g., "zeo")
	Money			int		// Captain's starting money
}

type SafeCaptain struct {
	ID 			string 	`json:"id"`
	Username 	string 	`json:"Username"`
	Rank 		string 	`json:"Rank"`
	Money 		int 	`json:"Money"`
}

// Define the CaptainRepository interface.
type CaptainRepository interface {
	Create(captain Captain) (string, error)
	Get(id string) (*Captain, error)
	GetByUsername(username string) (*Captain, error)
	Update(captain Captain) error
	Delete(id string) error
}

// sqliteCaptainRepository is a SQLite-based implementation of CaptainRepository.
type sqliteCaptainRepository struct {
	db *db.DB
}

var createTableQuery = `
CREATE TABLE IF NOT EXISTS captains (
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    rank TEXT NOT NULL,
    money INTEGER NOT NULL
);`

// NewCaptainRepository creates the captains table (if needed) and returns a repository.
func NewCaptainRepository(db *db.DB) (CaptainRepository, error) {
	// Step 1: Create the captains table using the SQL query.
	if err := db.CreateTable(createTableQuery); err != nil {
		return nil, err
	}

	// Step 2: Return a new instance of sqliteCaptainRepository
	return &sqliteCaptainRepository{db: db}, nil
}

// Create inserts a new captain record.
func (r *sqliteCaptainRepository) Create(captain Captain) (string, error) {
	// Step 1: Generate a new unique captain ID.
	captain.ID = cuid.New()

	// Step 2: Execute the INSERT query to add the captain record.
	query := `INSERT INTO captains (id, username, password_hash, rank, money) VALUES (?, ?, ?, ?, ?);`
	_, err := r.db.Exec(query, captain.ID, captain.Username, captain.PasswordHash, captain.Rank, captain.Money)
	if err != nil {
		return "", err
	}

	// Step 3: Return the new captain's ID.
	return captain.ID, nil
}

// Get retrieves a captain by ID.
func (r *sqliteCaptainRepository) Get(id string) (*Captain, error) {
	// Step 1: Execute a SELECT query to fetch the captain record by ID.
	query := `SELECT id, username, password_hash, rank, money FROM captains WHERE id = ?;`
	row := r.db.QueryRow(query, id)

	// Step 2: Scan the result into a Captain struct.
	var captain Captain
	err := row.Scan(&captain.ID, &captain.Username, &captain.PasswordHash, &captain.Rank, &captain.Money)

	// Step 3: Handle errors or absence of record.
	if err == sql.ErrNoRows {
		return nil, errors.New("captain not found")
	} else if err != nil {
		return nil, err
	}

	// Step 4: Return the captain record.
	return &captain, nil
}

// GetByUsername retrieves a captain by username.
func (r *sqliteCaptainRepository) GetByUsername(username string) (*Captain, error) {
	// Step 1: Execute a SELECT query to fetch the captain record by username.
	query := `SELECT id, username, password_hash, rank, money FROM captains WHERE username = ?;`
	row := r.db.QueryRow(query, username)

	// Step 2: Scan the result into a Captain struct.
	var captain Captain
	err := row.Scan(&captain.ID, &captain.Username, &captain.PasswordHash, &captain.Rank, &captain.Money)

	// Step 3: Handle errors or absence of record.
	if err == sql.ErrNoRows {
		return nil, errors.New("captain not found")
	} else if err != nil {
		return nil, err
	}

	// Step 4: Return the captain record.
	return &captain, nil
}

// Update modifies an existing captain record.
func (r *sqliteCaptainRepository) Update(captain Captain) error {
	// Step 1: Validate that the captain's ID is a valid CUID.
	if !isValidCUID(captain.ID) {
		return errors.New("invalid captain ID: not a valid CUID")
	}

	// Step 2: Execute the UPDATE query to modify the record.
	query := `UPDATE captains SET username = ?, password_hash = ?, rank = ?, money = ? WHERE id = ?;`
	result, err := r.db.Exec(query, captain.Username, captain.PasswordHash, captain.Rank, captain.Money, captain.ID)
	if err != nil {
		return err
	}

	// Step 3: Check the number of affected rows.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("captain not found")
	}

	// Step 4: Return nil indicating a successful update.
	return nil
}

// Delete removes a captain record by ID.
func (r *sqliteCaptainRepository) Delete(id string) error {
	// Step 1: Execute the DELETE query to remove the record.
	query := `DELETE FROM captains WHERE id = ?;`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	// Step 2: Check the number of affected rows.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("captain not found")
	}

	// Step 3: Return nil indicating successful deletion.
	return nil
}

// isValidCUID checks that the ID is 25 characters long and starts with 'c'.
func isValidCUID(id string) bool {

	// Step 1: Verify that the length is 25 and that it starts with 'c'.
	return len(id) == 25 && id[0] == 'c'
}