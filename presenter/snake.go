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
	return s.s.Occupies(x, y)
}

func (s *Snake) Length() int {
	return s.s.Length()
}
