package queue

// import (
// 	"sync"

// 	"github.com/tsraveling/dog-fight/server/internal/repositories"
// )

// // Match represents a single match instance.
// type Match struct {
// 	ID 			string
// 	GameMode 	GameMode
// 	Players 	[]*repositories.SafePlayer
// 	IsActive 	bool
// 	mu 			sync.Mutex
// }

// // NewMatch creates a new match with the given ID and mode.
// func NewMatch(id string, mode GameMode) *Match {
// 	return &Match{
// 		ID: id,
// 		GameMode: mode,
// 		Players: make([]*repositories.SafePlayer, 0, mode.MaxPlayers),
// 		IsActive: false,
// 	}
// }

// // AddPlayer adds a player to the match.
// func (m *Match) AddPlayer(player *repositories.SafePlayer) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()
// 	m.Players = append(m.Players, player)
// }

// // Start marks the match as active.
// func (m *Match) Start() {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()
// 	m.IsActive = true
// }

// // End marks the match as inactive.
// func (m *Match) End() {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()
// 	m.IsActive = false
// }