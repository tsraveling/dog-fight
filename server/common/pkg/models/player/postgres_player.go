package player

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
	"github.com/tsraveling/dog-fight/server/common/internal/database"
	"github.com/tsraveling/dog-fight/server/common/pkg/util"
)

// postgresPlayerRepository is a PostgreSQL-based implementation of PlayerRepository.
type postgresPlayerRepository struct {
	db *database.DB
}

// NewPostgresPlayerRepository creates the players table (if needed) and returns a PlayerRepository.
func NewPostgresPlayerRepository(database *database.DB) (PlayerRepository, error) {
	if err := database.CreateTable(createPlayerTableQuery); err != nil {
		return nil, err
	}
	return &postgresPlayerRepository{db: database}, nil
}

// Create inserts a new player record.
func (r *postgresPlayerRepository) Create(player Player) (string, error) {
	player.ID = cuid.New()
	query := `INSERT INTO players (id, username, password_hash, rank, money) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.db.Exec(query, player.ID, player.Username, player.PasswordHash, player.Rank, player.Money)
	if err != nil {
		return "", fmt.Errorf("failed to create player: %w", err)
	}
	return player.ID, nil
}

// Get retrieves a player by ID.
func (r *postgresPlayerRepository) Get(id string) (*Player, error) {
	query := `SELECT id, username, password_hash, rank, money FROM players WHERE id = $1;`
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
func (r *postgresPlayerRepository) GetByUsername(username string) (*Player, error) {
	query := `SELECT id, username, password_hash, rank, money FROM players WHERE username = $1;`
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
func (r *postgresPlayerRepository) Update(player Player) error {
	if !util.IsValidCUID(player.ID) {
		return errors.New("invalid player ID: not a valid CUID")
	}
	query := `UPDATE players SET username = $1, password_hash = $2, rank = $3, money = $4 WHERE id = $5;`
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
func (r *postgresPlayerRepository) Delete(id string) error {
	query := `DELETE FROM players WHERE id = $1;`
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
