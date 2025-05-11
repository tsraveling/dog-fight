package player

import (
	"database/sql"
	"errors"

	"github.com/lucsky/cuid"
	"github.com/tsraveling/dog-fight/server/common/internal/database"
	"github.com/tsraveling/dog-fight/server/common/pkg/util"
)

// sqlitePlayerRepository is a SQLite-based implementation of PlayerRepository.
type sqlitePlayerRepository struct {
	db *database.DB
}

// NewSQLitePlayerRepository creates the players table (if needed) and returns a PlayerRepository.
func NewSQLitePlayerRepository(database *database.DB) (PlayerRepository, error) {
	if err := database.CreateTable(createPlayerTableQuery); err != nil {
		return nil, err
	}
	return &sqlitePlayerRepository{db: database}, nil
}

// Create inserts a new player record.
func (r *sqlitePlayerRepository) Create(player Player) (string, error) {
	// Generate a new unique player ID.
	player.ID = cuid.New()
	query := `INSERT INTO players (id, username, password_hash, rank, money) VALUES (?, ?, ?, ?, ?);`
	_, err := r.db.Exec(query, player.ID, player.Username, player.PasswordHash, player.Rank, player.Money)
	if err != nil {
		return "", err
	}
	return player.ID, nil
}

// Get retrieves a player by ID.
func (r *sqlitePlayerRepository) Get(id string) (*Player, error) {
	query := `SELECT id, username, password_hash, rank, money FROM players WHERE id = ?;`
	row := r.db.QueryRow(query, id)
	var player Player
	err := row.Scan(&player.ID, &player.Username, &player.PasswordHash, &player.Rank, &player.Money)
	if err == sql.ErrNoRows {
		return nil, errors.New("player not found")
	} else if err != nil {
		return nil, err
	}
	return &player, nil
}

// GetByUsername retrieves a player by username.
func (r *sqlitePlayerRepository) GetByUsername(username string) (*Player, error) {
	query := `SELECT id, username, password_hash, rank, money FROM players WHERE username = ?;`
	row := r.db.QueryRow(query, username)
	var player Player
	err := row.Scan(&player.ID, &player.Username, &player.PasswordHash, &player.Rank, &player.Money)
	if err == sql.ErrNoRows {
		return nil, errors.New("player not found")
	} else if err != nil {
		return nil, err
	}
	return &player, nil
}

// Update modifies an existing player record.
func (r *sqlitePlayerRepository) Update(player Player) error {
	if !util.IsValidCUID(player.ID) {
		return errors.New("invalid player ID: not a valid CUID")
	}
	query := `UPDATE players SET username = ?, password_hash = ?, rank = ?, money = ? WHERE id = ?;`
	result, err := r.db.Exec(query, player.Username, player.PasswordHash, player.Rank, player.Money, player.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("player not found")
	}
	return nil
}

// Delete removes a player record by ID.
func (r *sqlitePlayerRepository) Delete(id string) error {
	query := `DELETE FROM players WHERE id = ?;`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("player not found")
	}
	return nil
}