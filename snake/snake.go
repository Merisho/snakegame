package snake

func NewSnake(x, y int) *Snake {
	head := newPart(x, y, NullVector)
	return &Snake{
		parts: []*Part{head},
		turns: [][3]int{},
	}
}

type Snake struct {
	parts []*Part
	turns [][3]int
}

func (s *Snake) Length() int {
	return len(s.parts)
}

func (s *Snake) Eat() {
	end := s.Tail().End()
	x := end.X()
	y := end.Y()

	switch end.Vector() {
	case DownVector:
		y--
	case UpVector:
		y++
	case LeftVector:
		x++
	case RightVector:
		x--
	case NullVector:
		x--
	}

	newPart := newPart(x, y, end.Vector())
	s.parts = append(s.parts, newPart)
}

func (s *Snake) X() int {
	return s.Head().X()
}

func (s *Snake) Y() int {
	return s.Head().Y()
}

func (s *Snake) Down() {
	s.Head().Down()
	s.pointTail()
}

func (s *Snake) Right() {
	s.Head().Right()
	s.pointTail()
}

func (s *Snake) Left() {
	s.Head().Left()
	s.pointTail()
}

func (s *Snake) Up() {
	s.Head().Up()
	s.pointTail()
}

func (s *Snake) pointTail() {
	if s.Tail().Start().Vector() == NullVector {
		// Because by default new snake parts are added to the left
		// So when it starts moving we must point the tail to the right
		s.Tail().SetVector(RightVector)
	}
}

func (s *Snake) Move() {
	s.Head().Move()
	prevVector := s.Head().Vector()
	if s.Length() == 1 {
		return
	}

	for _, part := range s.Tail().Parts() {
		part.Move()

		temp := part.Vector()
		part.setVector(prevVector)
		prevVector = temp
	}
}

func (s *Snake) Vector() int {
	return s.Head().Vector()
}

func (s *Snake) Tail() *Tail {
	if len(s.parts) == 1 {
		return &Tail{
			parts: s.parts[:],
		}
	}

	return &Tail{
		parts: s.parts[1:],
	}
}

func (s *Snake) Head() *Part {
	return s.parts[0]
}

func (s *Snake) CollidesItself() bool {
	for _, p := range s.parts[1:] {
		if s.X() == p.X() && s.Y() == p.Y() {
			return true
		}
	}

	return false
}
