package classes

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type hero struct {
	currentEnemy     *Enemy
	currentFightTime int
	enemiesDefeated  int
	heroType         heroType
	WaitGroup        sync.WaitGroup
	cancel           context.CancelFunc
	cancelContext    context.Context
}

type heroType string

const (
	moar heroType = "Moar"
	bulk          = "Bulk"
)

func newHero(t heroType) *hero {
	// create a cancel context to call
	// allows graceful exit and indefinite run
	ctx, cancel := context.WithCancel(context.Background())

	return &hero{
		cancel:        cancel,
		cancelContext: ctx,
		heroType:      t,
	}
}

// get human readable name based on type
func (h hero) GetHeroName() string {
	switch h.heroType {
	case moar:
		return "Moar, God of Static"
	case bulk:
		return "The Semi-Forgettable Bulk"
	default:
		return "Hero"
	}
}

// calls our hero to begin fighting an enemy
func (h *hero) FightEnemy(e *Enemy, c chan<- int) {
	// set current enemy
	h.currentEnemy = e

	// calculate fight time depending on parameters from prompt
	var fightTime int

	switch h.heroType {
	case moar:
		fightTime = calculateMoarFightTime(e)
	case bulk:
		fightTime = calculateBulkFightTime(e)
	}

	// start the fight
	h.WaitGroup.Add(1)
	go func() {
		defer h.WaitGroup.Done()
		h.startFight(fightTime, c)
	}()
}

// function to start a fight lasting n seconds
// which was determined in fight enemy function
func (h *hero) startFight(fightTime int, enemyDefeated chan<- int) {
	// create a timer to respond after n seconds
	fmt.Println(h.GetHeroName(), "starting a fight against", h.currentEnemy.GetEnemyName(), "lasting", fightTime, "seconds")
	fightTimer := time.NewTimer(time.Duration(fightTime) * time.Second)

	select {
	// cleanup when we get a cancel context
	case <-h.cancelContext.Done():
		fightTimer.Stop()
		return
	// respond to timer tick on channel
	case <-fightTimer.C:
		h.defeatEnemy(enemyDefeated)
	}
}

// calculate the fight time for Moar
func calculateMoarFightTime(e *Enemy) int {
	var fightTime int

	if e == nil {
		return fightTime
	}

	switch e.enemyType {
	// Moar + Alien: (powerLevel * 3) + random(1 or 2)
	case alien:
		fightTime += e.powerLevel*3 + rand.Intn(2) + 1
		break
	// Moar + Evil Robot: (powerLevel * 1) + random(1, 2, or 3)
	case robot:
		fightTime += e.powerLevel + rand.Intn(3) + 1
		break
	// Moar + Lizard Person: (powerLevel * 2) + random(1 or 2)
	case lizard:
		fightTime += e.powerLevel*2 + rand.Intn(2) + 1
		break
	}

	return fightTime
}

// calculate the fight time for Bulk
func calculateBulkFightTime(e *Enemy) int {
	var fightTime int

	if e == nil {
		return fightTime
	}

	switch e.enemyType {
	// Bulk + Alien: (powerLevel * 1) + random(0 or 2)
	case alien:
		// this is for the alien case - since the number is either zero or 2, but never 1.
		// if I misread the directions or they were mistyped, comment out this code and use
		// the following:

		//  delay += e.PowerLevel + rand.Intn(3)

		weirdRand := rand.Intn(2)

		if weirdRand == 1 {
			weirdRand += 1
		}

		fightTime += e.powerLevel + weirdRand
		break
	// Bulk + Evil Robot: (powerLevel * 3) + random(0 or 1)
	case robot:
		fightTime += e.powerLevel*3 + rand.Intn(2)
		break
	// Bulk + Lizard Person: (powerLevel * 2) + random(1 or 2)
	case lizard:
		fightTime += e.powerLevel*2 + rand.Intn(2) + 1
		break
	}

	return fightTime
}

// called when enough time has passed that we've defeated the enemy
func (h *hero) defeatEnemy(enemyDefeated chan<- int) {
	h.enemiesDefeated += 1

	fmt.Println(h.GetHeroName(), "has defeated", h.currentEnemy.GetEnemyName(), "total of", h.enemiesDefeated, "enemies defeated")

	h.currentEnemy = nil
	enemyDefeated <- 1
}

// returns whether or not we have an enemy
func (h hero) HasEnemy() bool {
	return h.currentEnemy != nil
}

// human readable representation of the hero along with current enemy (if any) and enemies defeated
func (h hero) String() string {
	if h.HasEnemy() {
		return fmt.Sprintf("%v has defeated %v enemies and is currently fighting %v", h.GetHeroName(), h.enemiesDefeated, h.currentEnemy.GetEnemyName())
	}

	return fmt.Sprintf("%v has no enemies, and has defeated %v enemies", h.GetHeroName(), h.enemiesDefeated)
}
