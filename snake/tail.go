package snake

type Tail struct {
	parts []*Part
}

func (t *Tail) SetVector(vector int) {
	for _, p := range t.parts {
		p.setVector(vector)
	}
}

func (t *Tail) Parts() []*Part {
	return t.parts
}

func (t *Tail) Start() *Part {
	return t.parts[0]
}

func (t *Tail) End() *Part {
	return t.parts[len(t.parts)-1]
}
