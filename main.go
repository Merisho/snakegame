package main

import (
	"github.com/merisho/snakegame/game"
	"github.com/merisho/snakegame/presenter"
	"github.com/merisho/snakegame/snake"
	"github.com/merisho/snakegame/view"
	"github.com/zetamatta/go-getch"
	"time"
)

const (
	up           = 119
	right        = 100
	down         = 115
	left         = 97
	exit         = 27
	mapWidth     = 20
	mapHeight    = 10
	tickInterval = 100 * time.Millisecond
)

func main() {
	s := snake.NewSnake(0, 0)
	snakePresenter := presenter.NewSnake(s)
	cliView := view.NewCLIView(snakePresenter, mapWidth, mapHeight)
	commandChan := make(chan rune)

	s.Down()

	g := game.NewGame(s, cliView, mapWidth, mapHeight)

	go func() {
		for {
			time.Sleep(tickInterval)
			g.Tick()
			if g.Over() {
				cliView.RenderStr("Game over! Press Escape to exit")
				return
			}
		}
	}()

	go func() {
		for {
			commandChan <- getch.Rune()
		}
	}()

	for commandCode := range commandChan {
		var command game.CommandFunc

		switch commandCode {
		case up:
			command = func(_ game.Snake) {
				s.Up()
			}
		case right:
			command = func(_ game.Snake) {
				s.Right()
			}
		case down:
			command = func(_ game.Snake) {
				s.Down()
			}
		case left:
			command = func(_ game.Snake) {
				s.Left()
			}
		case exit:
			return
		}

		g.Command(command)
	}
}
