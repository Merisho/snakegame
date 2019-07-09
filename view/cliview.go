package view

import (
	"fmt"
	"github.com/merisho/snakegame/presenter"
	"os"
	"os/exec"
	"runtime"
)

func NewCLIView(s *presenter.Snake, mapWidth, mapHeight int) *CLIView {
	return &CLIView{
		s:         s,
		mapWidth:  mapWidth,
		mapHeight: mapHeight,
	}
}

type CLIView struct {
	s         *presenter.Snake
	mapWidth  int
	mapHeight int
}

func (v *CLIView) Render(foodX, foodY int) {
	v.clear()

	fmt.Println("Score: ", v.s.Length())

	v.renderTopDownBorders()

	fmt.Print("\n")

	for i := 0; i < v.mapHeight; i++ {
		fmt.Print("|")
		for j := 0; j < v.mapWidth; j++ {
			if v.s.Occupies(j, i) {
				fmt.Print("o")
			} else if i == foodY && j == foodX {
				fmt.Print("☻")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}

	v.renderTopDownBorders()

	fmt.Print("\n")
}

func (v *CLIView) RenderStr(s string) {
	fmt.Println(s)
}

func (v *CLIView) clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func (v *CLIView) renderTopDownBorders() {
	fmt.Print("|")
	for i := 0; i < v.mapWidth; i++ {
		fmt.Print("—")
	}
	fmt.Print("|")
}
