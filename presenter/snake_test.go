package presenter

import (
	"github.com/merisho/snakegame/snake"
	"testing"
)

func TestOccupies(t *testing.T) {
	s := snake.NewSnake(0, 0)
	s.Eat()
	presenter := Snake{s: s}

	if !presenter.Occupies(s.X(), s.Y()) {
		t.Fatal("Must return true if snake's head occupies given coordinates")
	}

	tail := s.Tail().Start()
	if !presenter.Occupies(tail.X(), tail.Y()) {
		t.Fatal("Must return true if snake's tail occupies given coordinates")
	}

	if presenter.Occupies(100, 100) {
		t.Fatal("Must return false if snake does not occupy given coordinates")
	}
}
