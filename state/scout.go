package state

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type scout struct {
	waitGroup     sync.WaitGroup
	cancel        context.CancelFunc
	cancelContext context.Context
}

func newScout() *scout {
	// create a cancel context to call
	// allows graceful exit and indefinite run
	ctx, cancel := context.WithCancel(context.Background())

	return &scout{
		cancel:        cancel,
		cancelContext: ctx,
	}
}

// create a scout state where after two seconds we spawn
// n random enemies with timeouts; between 1 and 3 enemies
func (s *scout) scout(currentState *GameState) {
	// create a timer to respond after 2 seconds
	fmt.Println("EagleWoman scouting for new targets (takes 2 seconds)")
	scoutTimer := time.NewTimer(2 * time.Second)

	select {
	// cleanup when we get a cancel context
	case <-s.cancelContext.Done():
		scoutTimer.Stop()
		return
	// respond to timer tick on channel
	case <-scoutTimer.C:
		// generate a random between 1 and 3
		randomNum := rand.Intn(3) + 1
		fmt.Println("EagleWoman scouted", randomNum, "new target(s)")

		// spawn n number of enemies between 1-3
		for i := 0; i < randomNum; i++ {
			currentState.WaitGroup.Add(1)
			go func() {
				defer currentState.WaitGroup.Done()
				currentState.SpawnEnemy()
			}()
		}
	}
}
