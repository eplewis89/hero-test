package classes

import (
	"fmt"
	"sync"
)

type Team struct {
	moar      *hero
	bulk      *hero
	waitGroup sync.WaitGroup
}

func NewTeam() *Team {
	return &Team{
		moar: newHero(moar),
		bulk: newHero(bulk),
	}
}

func (t *Team) GetMoar() *hero {
	return t.moar
}

func (t *Team) GetBulk() *hero {
	return t.bulk
}

// human readable string representation of team
func (t Team) String() string {
	return fmt.Sprintf("%v\n%v", t.moar, t.bulk)
}

func (t Team) GetTotalDefeated() int {
	return t.bulk.enemiesDefeated + t.moar.enemiesDefeated
}
