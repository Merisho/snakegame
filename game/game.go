package game

func NewGame(s Snake, v View, maxX, maxY int) *Game {
	foodX, foodY := generateNewFood(maxX, maxY)
	return &Game{
		foodX: foodX,
		foodY: foodY,
		s:     s,
		v:     v,
		maxX:  maxX,
		maxY:  maxY,
	}
}

type Game struct {
	foodX int
	foodY int
	s     Snake
	v     View
	comm  CommandFunc
	over  bool
	maxX  int
	maxY  int
}

type CommandFunc func(s Snake)

func (g *Game) Over() bool {
	return g.over
}

func (g *Game) Tick() {
	if g.comm != nil {
		g.comm(g.s)
		g.comm = nil
	}

	g.s.Move()

	if g.snakeBumped() {
		g.over = true
		return
	} else if g.snakeCollidedFood() {
		g.s.Eat()
		g.foodX, g.foodY = generateNewFood(g.maxX, g.maxY)
	}

	g.v.Render(g.foodX, g.foodY)
}

func (g *Game) snakeCollidedFood() bool {
	return g.s.X() == g.foodX && g.s.Y() == g.foodY
}

func (g *Game) snakeBumped() bool {
	bumpedBorder := g.s.X() == g.maxX || g.s.Y() == g.maxY || g.s.X() == -1 || g.s.Y() == -1
	return bumpedBorder || g.s.CollidesItself()
}

func (g *Game) Command(f CommandFunc) {
	g.comm = f
}
