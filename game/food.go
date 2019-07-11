package game

import (
	"math/rand"
	"time"
)

var foodRnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var generateNewFood = func(maxX, maxY int, s Snake) (x, y int) {
	x, y = foodRnd.Intn(maxX), foodRnd.Intn(maxY)

	for s.Occupies(x, y) {
		x, y = foodRnd.Intn(maxX), foodRnd.Intn(maxY)
	}

	return
}
