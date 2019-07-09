package snake

import (
	"testing"
)

func TestStartingLength(t *testing.T) {
	snake := NewSnake(0, 0)

	if snake.Length() != 1 {
		t.Fatal("Starting snake length must be 1")
	}
}

func TestEat(t *testing.T) {
	snake := NewSnake(0, 0)

	snake.Eat()
	if snake.Length() != 2 {
		t.Fatal("Snake length must be increased after snake ate")
	}
}

func TestStartingPosition(t *testing.T) {
	snake := NewSnake(10, 10)

	if snake.X() != 10 || snake.Y() != 10 {
		t.Fatal("Must set starting position")
	}
}

func TestStartingTail(t *testing.T) {
	snake := NewSnake(0, 0)
	if snake.Tail().End().X() != 0 || snake.Tail().End().Y() != 0 {
		t.Fatal("Starting tail coords must be same as head")
	}
}

func TestHeadBodyPart(t *testing.T) {
	snake := NewSnake(0, 0)

	head := snake.Head()
	if head.X() != snake.X() || head.Y() != snake.Y() {
		t.Fatal("First part must be head with snake coordinates")
	}
}

func TestChangeSnakeVector(t *testing.T) {
	snake := NewSnake(10, 10)

	snake.Eat()
	snake.Up()

	if snake.Vector() != UpVector {
		t.Fatal("Must change head vector")
	}

	if snake.Tail().End().Vector() != RightVector {
		t.Fatal(`Must point tail to the RIGHT by default
						since all new snake parts are added to the left by default`)
	}
}

func TestSnakeNullVectorNewPartPotision(t *testing.T) {
	snake := NewSnake(0, 0)

	oldTail := snake.Tail()
	snake.Eat()
	if snake.Tail().End().X() != oldTail.End().X()-1 {
		t.Fatal("If the snake has Null vector newly appended part must be added to the left")
	}
}

func TestSnakeNewPartPosition(t *testing.T) {
	snake := NewSnake(10, 10)

	snake.Left()
	oldTail := snake.Tail()
	snake.Eat()
	if snake.Tail().End().X() != oldTail.End().X()+1 {
		t.Fatal("Must append new part to coordinate with opposite vector")
	}

	if snake.Tail().End().Vector() != oldTail.End().Vector() {
		t.Fatal("Newly appended tail must have the same vector as previous tail")
	}
}

func TestSnakeMovementSinglePart(t *testing.T) {
	snake := NewSnake(0, 0)

	snake.Down()
	snake.Move()
	snake.Right()
	snake.Move()

	if snake.X() != 1 {
		t.Fatal("Must move snake for 1 point only by X")
	}

	if snake.Y() != 1 {
		t.Fatal("Must move snake for 1 point only by Y")
	}
}

func TestSnakeMovement(t *testing.T) {
	snake := NewSnake(10, 10)
	snake.Eat()

	snake.Down()
	snake.Move()

	if snake.Y() != 11 {
		t.Fatal("Must move snake head to the bottom")
	}

	if snake.Tail().Start().X() != 10 {
		t.Fatal("Must move tail to the previous head position")
	}
}

func TestCollidesItself(t *testing.T) {
	snake := NewSnake(10, 10)
	snake.parts = []*Part{
		{ Point: &Point{x: 10, y: 10} },
		{ Point: &Point{x: 11, y: 10} },
		{ Point: &Point{x: 11, y: 9} },
		{ Point: &Point{x: 10, y: 10} },
	}

	if snake.CollidesItself() {
		return
	}

	t.Fatal("Must return true in case snake collides itself")
}
