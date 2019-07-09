package presenter

import "github.com/merisho/snakegame/snake"

func NewSnake(s *snake.Snake) *Snake {
	return &Snake{
		s: s,
	}
}

type Snake struct {
	s *snake.Snake
}

func (s *Snake) Occupies(x, y int) bool {
	if s.s.Head().X() == x && s.s.Head().Y() == y {
		return true
	}

	for _, part := range s.s.Tail().Parts() {
		if part.X() == x && part.Y() == y {
			return true
		}
	}

	return false
}

func (s *Snake) Length() int {
	return s.s.Length()
}
