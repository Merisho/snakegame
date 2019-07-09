package game

type SpySnake struct {
	eatCallCount  int
	moveCallCount int
	upCallCount   int
	x             int
	y             int
}

func (s *SpySnake) X() int {
	return s.x
}

func (s *SpySnake) Y() int {
	return s.y
}

func (s *SpySnake) Move() {
	s.moveCallCount++
}

func (s *SpySnake) Eat() {
	s.eatCallCount++
}

func (s *SpySnake) CollidesItself() bool {
	return false
}

type SpyView struct {
	renderCallCount int
}

func (v *SpyView) Render(foodX, foodY int) {
	v.renderCallCount++
}
