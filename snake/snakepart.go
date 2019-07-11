package snake

const (
	NullVector = iota
	DownVector
	RightVector
	UpVector
	LeftVector
)

func newPart(x, y, vector int) *Part {
	return &Part{
		x: x,
		y: y,
		vector: vector,
	}
}

type Part struct {
	x int
	y int
	vector int
}

func (p *Part) X() int {
	return p.x
}

func (p *Part) Y() int {
	return p.y
}

func (p *Part) Vector() int {
	return p.vector
}

func (p *Part) Move() {
	switch p.Vector() {
	case DownVector:
		p.y++
	case UpVector:
		p.y--
	case RightVector:
		p.x++
	case LeftVector:
		p.x--
	}
}

func (p *Part) Down() {
	p.setVector(DownVector)
}

func (p *Part) Right() {
	p.setVector(RightVector)
}

func (p *Part) Up() {
	p.setVector(UpVector)
}

func (p *Part) Left() {
	p.setVector(LeftVector)
}

func (p *Part) setVector(vector int) {
	if vector == DownVector && p.Vector() != UpVector {
		p.vector = vector
	} else if vector == UpVector && p.Vector() != DownVector {
		p.vector = vector
	} else if vector == RightVector && p.Vector() != LeftVector {
		p.vector = vector
	} else if vector == LeftVector && p.Vector() != RightVector {
		p.vector = vector
	}
}
