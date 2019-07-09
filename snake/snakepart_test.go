package snake

import "testing"

func TestChangeVector(t *testing.T) {
	part := &Part{}

	part.Down()
	if part.Vector() != DownVector {
		t.Fatal("Must change direction to down")
	}

	part.Right()
	if part.Vector() != RightVector {
		t.Fatal("Must change direction to right")
	}

	part.Up()
	if part.Vector() != UpVector {
		t.Fatal("Must change direction to up")
	}

	part.Left()
	if part.Vector() != LeftVector {
		t.Fatal("Must change direction to left")
	}
}

func TestChangeToOppositeVector(t *testing.T) {
	part := &Part{}

	part.Up()
	part.Down()
	if part.Vector() != UpVector {
		t.Fatal("Must not change the vector if requested vector is totally opposite to the current")
	}

	part.Right()
	part.Left()
	if part.Vector() != RightVector {
		t.Fatal("Must not change the vector if requested vector is totally opposite to the current")
	}
}

func TestMove(t *testing.T) {
	part := newPart(0, 0, 0)

	part.Down()
	part.Move()
	if part.Y() != 1 {
		t.Fatal("Down direction move must increment Y coordinate")
	}

	part.Right()
	part.Move()
	if part.X() != 1 {
		t.Fatal("Right direction move must increment X coordinate")
	}

	part.Up()
	part.Move()
	if part.Y() != 0 {
		t.Fatal("Up direction move must decrement Y coordinate")
	}

	part.Left()
	part.Move()
	if part.X() != 0 {
		t.Fatal("Left direction move must decrement X coordinate")
	}
}
