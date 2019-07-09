package game

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var generateNewFood = func(maxX, maxY int) (x, y int) {
	return rnd.Intn(maxX), rnd.Intn(maxY)
}
