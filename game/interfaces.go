package game

type Snake interface {
	Move()
	Eat()
	CollidesItself() bool
	Occupies(x, y int) bool
	X() int
	Y() int
}

type View interface {
	Render(foodX, foodY int)
}
