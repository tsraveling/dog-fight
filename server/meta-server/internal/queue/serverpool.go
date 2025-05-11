package queue

// import (
// 	"errors"
// 	"sync"

// 	"github.com/tsraveling/dog-fight/server/internal/repositories"
// )

// // ServerPool manages multiple QueueManager instances (each a separate server).
// type ServerPool struct {
// 	mu 			sync.Mutex
// 	queues 		[]*QueueManager	// Active server queues
// 	maxServers 	int				// Global config: maximum number of servers allowed
// 	gameMode 	GameMode		// The game mode these servers run in
// }

// // NewServerPool creates a new ServerPool for a given game mode and maximum server count.
// func NewServerPool(gameMode GameMode, maxServers int) *ServerPool {
// 	return &ServerPool{
// 		queues: make([]*QueueManager, 0),
// 		maxServers: maxServers,
// 		gameMode: gameMode,
// 	}
// }

// // AddPlayer assigns a player to an existing queue if there is room; if not,
// // and if fewer than maxServers exists, it creates a new QueueManager. If all queues are full,
// // it distributes the new Player to the queue with the fewest waiting players.
// func (sp *ServerPool) AddPlayer(player *repositories.SafePlayer) error {
// 	sp.mu.Lock()
// 	defer sp.mu.Unlock()

// 	var target *QueueManager
// 	minCount := int(^uint(0) >> 1) // Initialize with maximum int value

// 	// Try to find a queue with room (waiting count less than MaxPlayers).
// 	for _, qm := range sp.queues {
// 		count := len(qm.waiting)
// 		if count < sp.gameMode.MaxPlayers {
// 			if count < minCount {
// 				minCount = count
// 				target = qm
// 			}
// 		}
// 	}

// 	// If no queue has room and we haven't reached the max server limit, create a new queue.
// 	if target == nil {
// 		if len(sp.queues) < sp.maxServers {
// 			newQM := NewQueueManager(sp.gameMode)
// 			sp.queues = append(sp.queues, newQM)
// 			target = newQM
// 		} else {
// 			// All queues are full; choose the one with the fewest waiting players.
// 			for _, qm := range sp.queues {
// 				count := len(qm.waiting)
// 				if count < minCount {
// 					minCount = count
// 					target = qm
// 				}
// 			}
// 		}
// 	}

// 	if target == nil {
// 		return errors.New("no available server queue")
// 	}

// 	// Add the player to the selected queue.
// 	target.AddPlayer(player)
// 	return nil
// }

// // GetQueues returns a copy of all current QueueManagers.
// func (sp *ServerPool) GetQueues() []*QueueManager {
// 	sp.mu.Lock()
// 	defer sp.mu.Unlock()

// 	copyQueues := make([]*QueueManager, len(sp.queues))
// 	copy(copyQueues, sp.queues)
// 	return copyQueues
// }