package queue

import (
	"errors"
	"sync"

	"github.com/tsraveling/dog-fight/server/internal/repositories"
)

// MatchQueue manages the queue for one server.
type MatchQueue struct {
	mu 				sync.Mutex
	ServerID 		string
	activePlayers	[]repositories.Captain
	nextLobby 		[]repositories.Captain
	waitingQueue 	[]repositories.Captain
	maxActive 		int
	maxLobby 		int
}

// NewMatchQueue creates a new MatchQueue for a given server.
func NewMatchQueue(serverID string, maxActive, maxLobby int) *MatchQueue {
	return &MatchQueue{
		ServerID: serverID,
		maxActive: maxActive,
		maxLobby: maxLobby,
		activePlayers: make([]repositories.Captain, 0, maxActive),
		nextLobby: make([]repositories.Captain, 0, maxLobby),
		waitingQueue: make([]repositories.Captain, 0),
	}
}

// Join adds a captain to the waiting queue and promites players if possible.
func (mq *MatchQueue) Join(captain repositories.Captain) {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	mq.waitingQueue = append(mq.waitingQueue, captain)
	mq.promote()
}

// promote moves captains from waitingQueue to nextLobby until nextLobby is full.
func (mq *MatchQueue) promote() {
	for len(mq.nextLobby) < mq.maxLobby && len(mq.waitingQueue) > 0 {
		captain := mq.waitingQueue[0]
		mq.waitingQueue = mq.waitingQueue[1:]
		mq.nextLobby = append(mq.nextLobby, captain)
	}
}

// StartMatch promotes captains from the next lobby to active play.
// It returns an error if a match is already active or if there are not enough players.
func (mq *MatchQueue) StartMatch() error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if len(mq.activePlayers) > 0 {
		return errors.New("match already in progress")
	}
	if len(mq.nextLobby) < mq.maxActive {
		return errors.New("not enough players in next lobby to start match")
	}

	// Promote the first maxActive captains from nextLobby into activePlayers.
	mq.activePlayers = append(mq.activePlayers, mq.nextLobby[:mq.maxActive]...)
	mq.nextLobby = mq.nextLobby[mq.maxActive:]
	mq.promote() // Refill nextLobby from waitingQueue, if possible.
	return nil
}

// EndMatch clears the active match (activePlayers) and promotes waiting players.
func (mq *MatchQueue) EndMatch() {
	mq.mu.Lock()
	defer mq.mu.Unlock()
	mq.activePlayers = mq.activePlayers[:0]
	mq.promote()
}

// GetStatus returns copies of the current active, lobby, and waiting lists.
func (mq *MatchQueue) GetStatus() (active, lobby, waiting []repositories.Captain) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	active = make([]repositories.Captain, len(mq.activePlayers))
	copy(active, mq.activePlayers)
	lobby = make([]repositories.Captain, len(mq.nextLobby))
	copy(lobby, mq.nextLobby)
	waiting = make([]repositories.Captain, len(mq.waitingQueue))
	copy(waiting, mq.waitingQueue)
	return
}