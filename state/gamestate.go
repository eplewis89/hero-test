package state

import (
	"container/list"
	"context"
	"fmt"
	"math/rand"
	"strateos-test/classes"
	"sync"
	"time"
)

type GameState struct {
	defeated       int
	runtime        int
	WaitGroup      sync.WaitGroup
	cancel         context.CancelFunc
	cancelContext  context.Context
	enemyListMutex sync.Mutex
	enemyAdded     chan int
	enemyDefeated  chan int
	frontLine      *list.List
	heroTeam       *classes.Team
	eagleWoman     *scout
}

// create a new game state object to keep track of
// the current run state
func NewGameState() *GameState {
	// create a cancel context to call
	// allows graceful exit and indefinite run
	ctx, cancel := context.WithCancel(context.Background())

	return &GameState{
		cancel:        cancel,
		cancelContext: ctx,
		enemyAdded:    make(chan int, 100),
		enemyDefeated: make(chan int, 100),
		frontLine:     list.New(),
		heroTeam:      classes.NewTeam(),
		eagleWoman:    newScout(),
	}
}

// increments the game runtime
func (s *GameState) incrementRuntime() {
	s.runtime += 1
}

// creates n number of enemies
func (s *GameState) SpawnEnemy() {
	// create a delay from 1-5 seconds
	delay := rand.Intn(5) + 1
	enemyTimer := time.NewTimer(time.Duration(delay) * time.Second)
	fmt.Println("Enemy takes", delay, "seconds to spot")

	select {
	// cleanup
	case <-s.cancelContext.Done():
		return
	// respond to timer tick
	case <-enemyTimer.C:
		fmt.Println("Enemy spotted after", delay, "seconds")

		// create an enemy
		spawned := classes.GetRandomEnemy(delay)

		s.addEnemy(spawned)
	}
}

// add an enemy to the back of the list
func (s *GameState) addEnemy(e *classes.Enemy) {
	s.enemyListMutex.Lock()
	defer s.enemyListMutex.Unlock()

	s.frontLine.PushBack(e)

	s.enemyAdded <- 1
}

// get count of list of enemies on front line
func (s *GameState) hasNextEnemy() bool {
	s.enemyListMutex.Lock()
	defer s.enemyListMutex.Unlock()

	return s.frontLine.Len() > 0
}

// grab the next enemy off the front of the list
func (s *GameState) getNextEnemy() *classes.Enemy {
	s.enemyListMutex.Lock()
	defer s.enemyListMutex.Unlock()

	// check if we have a value on the front of the list
	// to get off of the list
	if item := s.frontLine.Front(); item != nil {
		// grab the enemy value while also removing it
		return s.frontLine.Remove(item).(*classes.Enemy)
	}

	// if we have no enemies to return just give nil
	return nil
}

// using the observer pattern we can determine when
// an enemy has been on the front line for too long
func (s *GameState) RunObserver() {
	ticker := time.NewTicker(1 * time.Second)

	// since this is a static 1 second observable
	// use a ticker - this will be easier than rescheduling
	// the observe task as in the other methods
	// also run indefinitely
	for {
		select {
		// in the case we receive a quit message on the subbed channel
		// stop the ticker and return to cleanly exit the thread
		case <-s.cancelContext.Done():
			ticker.Stop()
			return
		// when we have a tick...
		case <-ticker.C:
			// keep track of running time
			s.incrementRuntime()

			// show runtime
			// fmt.Println("(observing front line...", s.Runtime, "seconds)")

			// iterate through the frontline and check
			// if any enemies have been waiting for over 10 seconds
			for e := s.frontLine.Front(); e != nil; e = e.Next() {
				cur := e.Value.(*classes.Enemy)

				// if so end the program because at this point the battle is lost
				if cur.GetOnFieldTime() >= 10 {
					fmt.Println("Enemy has breached the frontlines, destroying our servers in the process")
					fmt.Println("Final run time:", s.runtime)
					fmt.Println(s.heroTeam)
					ticker.Stop()
					s.stopGame()
					return
				}

				cur.IncrementOnFieldTime()

				fmt.Println("{ ", cur, " }")
			}
		}
	}
}

// wait for n number of seconds from 5-10
// then begin the notification of new enemies appearing on
// the frontline
func (s *GameState) RunArachnidSense() {
	for {
		// create a timer for n seconds in the future from 5-10
		randomInterval := rand.Intn(6) + 5
		spideyTimer := time.NewTimer(time.Duration(randomInterval) * time.Second)
		fmt.Println("There's something on the horizon... waiting for", randomInterval, "seconds")

		select {
		case <-s.cancelContext.Done():
			spideyTimer.Stop()
			return
		case <-spideyTimer.C:
			fmt.Println("Threat detected after", randomInterval, "seconds")
			s.eagleWoman.scout(s)
		}
	}
}

// starts a protection routine which checks the list of enemies to dispatch
func (s *GameState) RunProtectServers() {
	for {
		select {
		case <-s.cancelContext.Done():
			return
		case <-s.enemyDefeated:
			s.nofityEnemyAvailable()
		case <-s.enemyAdded:
			s.nofityEnemyAvailable()
		}
	}
}

func (s *GameState) nofityEnemyAvailable() {
	if !s.heroTeam.GetMoar().HasEnemy() && s.hasNextEnemy() {
		e := s.getNextEnemy()
		s.heroTeam.GetMoar().FightEnemy(e, s.enemyDefeated)
	}

	if !s.heroTeam.GetBulk().HasEnemy() && s.hasNextEnemy() {
		e := s.getNextEnemy()
		s.heroTeam.GetBulk().FightEnemy(e, s.enemyDefeated)
	}
}

// called when a stopgame has been initiated
func (s *GameState) stopGame() {
	s.eagleWoman.cancel()
	s.cancel()
}

func (s *GameState) GetTotalDefeated() int {
	return s.heroTeam.GetTotalDefeated()
}
