package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	var genFoodCalled bool
	generateNewFood = func(maxX, maxY int) (foodX, foodY int) {
		genFoodCalled = true
		return 0, 0
	}
	NewGame(&SpySnake{}, &SpyView{}, 1, 1)

	if genFoodCalled {
		return
	}

	t.Fatal("Must generate a piece of food on game start")
}

func TestGameTickSnakeMovement(t *testing.T) {
	snake := &SpySnake{}
	game := NewGame(snake, &SpyView{}, 1, 1)

	game.Tick()

	if snake.moveCallCount == 1 {
		return
	}

	t.Fatal("Must move snake on game tick")
}

func TestGameTickFoodEating(t *testing.T) {
	s := &SpySnake{
		x: 10,
		y: 10,
	}
	generateNewFood = func(maxX, maxY int) (foodX, foodY int) {
		return 10, 10
	}
	game := NewGame(s, &SpyView{}, 100, 100)

	game.Tick()

	if s.eatCallCount == 1 {
		return
	}

	t.Fatal("Must eat the food in case snake encountered one")
}

func TestGameTickFoodGeneration(t *testing.T) {
	s := &SpySnake{
		x: 10,
		y: 10,
	}
	var genFoodCallCount int
	generateNewFood = func(maxX, maxY int) (foodX, foodY int) {
		genFoodCallCount++
		return 10, 10
	}
	game := NewGame(s, &SpyView{}, 100, 100)

	game.Tick()

	if genFoodCallCount == 2 {
		return
	}

	t.Fatal("Must generate a new piece of food when snake has eaten one")
}

func TestGameTickRender(t *testing.T) {
	view := &SpyView{}
	game := NewGame(&SpySnake{}, view, 1, 1)

	game.Tick()

	if view.renderCallCount == 1 {
		return
	}

	t.Fatal("Must render the view on game tick")
}

func TestGameTickCommand(t *testing.T) {
	var commandExecuted bool
	game := NewGame(&SpySnake{}, &SpyView{}, 1, 1)

	game.Command(func(s Snake) {
		commandExecuted = true
	})
	game.Tick()

	if commandExecuted {
		return
	}

	t.Fatal("Must execute command on game tick")
}

func TestGameOver(t *testing.T) {
	maxX, maxY := 1, 1
	testCases := []struct {
		x int
		y int
	}{
		{
			x: -1,
			y: 0,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: maxX,
			y: 0,
		},
		{
			x: 0,
			y: maxY,
		},
	}

	for _, c := range testCases {
		s := &SpySnake{
			x: c.x,
			y: c.y,
		}
		game := NewGame(s, &SpyView{}, maxX, maxY)

		game.Tick()

		if game.Over() {
			continue
		}

		t.Fatal("If snake collides the border the game must be over")
	}
}
