package queue

// import (
// 	"errors"
// 	"sync"

// 	"github.com/lucsky/cuid"
// 	"github.com/tsraveling/dog-fight/server/internal/repositories"
// )

// // QueueManager manages waiting players and creates matches.
// type QueueManager struct {
// 	mu 			sync.Mutex
// 	waiting 	[]*repositories.SafePlayer
// 	gameMode 	GameMode
// 	activeMatch	*Match
// }

// // NewQueueManager creates a new QueueManager for the given mode.
// func NewQueueManager(mode GameMode) *QueueManager {
// 	return &QueueManager{
// 		waiting: make([]*repositories.SafePlayer, 0),
// 		gameMode: mode,
// 	}
// }

// // AddPlayer appends a new player to the waiting list, then tries to start a match.
// func (qm *QueueManager) AddPlayer(player *repositories.SafePlayer) {
// 	qm.mu.Lock()
// 	defer qm.mu.Unlock()

// 	// Step 1: Add to waiting list.
// 	qm.waiting = append(qm.waiting, player)

// 	// Step 2: Check if we can start a match.
// 	qm.maybeStartMatch()
// }

// // maybeStartmatch checks if enough players are waiting to start a match.
// func (qm *QueueManager) maybeStartMatch() {
// 	if len(qm.waiting) < qm.gameMode.MaxPlayers {
// 		return
// 	}

// 	// Step 1: Create a new match once we have enough players.
// 	matchID := cuid.New()
// 	match := NewMatch(matchID, qm.gameMode)

// 	// Step 2: Move exactly MaxPlayers from waiting to the match.
// 	for i := 0; i < qm.gameMode.MaxPlayers; i++ {
// 		player := qm.waiting[0]
// 		qm.waiting = qm.waiting[1:]
// 		match.AddPlayer(player)
// 	}

// 	// Step 3: Mark the match as active.
// 	match.Start()
// 	qm.activeMatch = match
// }

// // GetActiveMatch returns the currently active match (or nil if none).
// func (qm *QueueManager) GetActiveMatch() *Match {
// 	qm.mu.Lock()
// 	defer qm.mu.Unlock()
// 	return qm.activeMatch
// }

// // EndMatch ends the current active match, if one exists.
// func (qm *QueueManager) EndMatch() error {
// 	qm.mu.Lock()
// 	defer qm.mu.Unlock()

// 	if qm.activeMatch == nil || !qm.activeMatch.IsActive {
// 		return errors.New("no active match to end")
// 	}

// 	qm.activeMatch.End()
// 	qm.activeMatch = nil
// 	return nil
// }