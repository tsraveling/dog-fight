package queue

import (
	"errors"
	"sync"

	"github.com/tsraveling/dog-fight/server/internal/repositories"
)

// ServerPool manages one or more MatchQueue instances.
type ServerPool struct {
	mu 		sync.Mutex
	queues 	map[string]*MatchQueue
}

// NewServerPool creates a new, empty ServerPool.
func NewServerPool() *ServerPool {
	return &ServerPool{
		queues: make(map[string]*MatchQueue),
	}
}

// AddQueue adds a new MatchQueue to the pool.
func (sp *ServerPool) AddQueue(mq *MatchQueue) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	sp.queues[mq.ServerID] = mq
}

// GetQueue retrieves a MatchQueue by server ID.
func (sp *ServerPool) GetQueue(serverID string) (*MatchQueue, error) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	if queue, ok := sp.queues[serverID]; ok {
		return queue, nil
	}
	return nil, errors.New("server not found")
}

// AddCaptain routes a captain to the least loaded server.
func (sp *ServerPool) AddCaptain(captain repositories.Captain) error {
	sp.mu.Lock()
	defer sp.mu.Unlock()

	if len(sp.queues) == 0 {
		return errors.New("no servers available")
	}

	var selected *MatchQueue
	minLoad := -1

	// Select the server with the lowest total load.
	for _, mq := range sp.queues {
		mq.mu.Lock()
		load := len(mq.activePlayers) + len(mq.nextLobby) + len(mq.waitingQueue)
		mq.mu.Unlock()

		if selected == nil || load < minLoad || minLoad == -1 {
			selected = mq
			minLoad = load
		}
	}

	if selected == nil {
		return errors.New("could not select a server")
	}

	selected.Join(captain)
	return nil
}