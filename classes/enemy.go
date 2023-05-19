package classes

import (
	"fmt"
	"math/rand"
)

type Enemy struct {
	enemyType    enemyType
	powerLevel   int
	timeToArrive int
	onFieldTime  int
}

type enemyType int

const (
	alien  enemyType = iota
	robot            = iota
	lizard           = iota
)

// generate a random enemy based on params from prompt
func GetRandomEnemy(delay int) *Enemy {
	return &Enemy{
		enemyType:    enemyType(rand.Intn(3)),
		powerLevel:   rand.Intn(3) + 1,
		timeToArrive: delay,
	}
}

// human readable enemy name
func (e Enemy) GetEnemyName() string {
	switch e.enemyType {
	case alien:
		return "Alien"
	case robot:
		return "Evil Robot"
	case lizard:
		return "Lizard Person"
	default:
		return "Enemy"
	}
}

// increments on field time when enemy is on the field
func (e *Enemy) IncrementOnFieldTime() {
	e.onFieldTime += 1
}

func (e Enemy) GetOnFieldTime() int {
	return e.onFieldTime
}

// human readable string representation
func (e Enemy) String() string {
	return fmt.Sprintf("%v (powerLevel %v, timeToArrive %v) on field for %v seconds", e.GetEnemyName(), e.powerLevel, e.timeToArrive, e.onFieldTime)
}
