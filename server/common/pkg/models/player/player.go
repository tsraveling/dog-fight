package player

// Player represents a user in the system.
type Player struct {
	ID 				string
	Username 		string
	PasswordHash 	string
	Rank 			string
	Money 			int
}

// SafePlayer is a sanitized version of Player for responses.
type SafePlayer struct {
	ID 			string `json:"id"`
	Username 	string `json:"username"`
	Rank 		string `json:"rank"`
	Money 		string `json:"money"`
}

// PlayerRepository defines operations for managing Player records.
type PlayerRepository interface {
	Create(player Player) (string, error)
	Get(id string) (*Player, error)
	GetByUsername(username string) (*Player, error)
	Update(player Player) error
	Delete(id string) error
}

var createPlayerTableQuery = `
CREATE TABLE IF NOT EXISTS players (
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    rank TEXT NOT NULL,
    money INTEGER NOT NULL
);`
