package board

import (
	"github.com/hazge/go-snake/cursor"
	"github.com/hazge/go-snake/gu"
	"github.com/hazge/go-snake/snake"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Board struct {
	YLength uint16
	XLength uint16
	Cursor  cursor.Cursor
	Snake   snake.Snake
	y       string
}

func New() Board {
	w, h := getTerminalSize()
	b := Board{w, h, cursor.New(), snake.New(w, h), "a"}
	b.CleanBoard()
	return b

}

func getTerminalSize() (uint16, uint16) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	size := strings.Split(string(out), " ")
	w, err := strconv.ParseInt(size[0], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	h, err := strconv.ParseInt(strings.TrimRight(size[1], "\n"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	if w <= 0 || h <= 0 {
		panic("empty terminal")
	}
	return uint16(w), uint16(h)
}
func (b *Board) DrawBorder() {
	gu.RunCommand("[42m")
	// left border
	for b.Cursor.Position.Y < b.YLength {
		b.Cursor.DrawAndBack(" ")
		b.Cursor.MoveDown()
	}

	for b.Cursor.Position.X < b.XLength {
		b.Cursor.DrawAndBack(" ")
		b.Cursor.MoveRight()
	}
	for b.Cursor.Position.Y > 0 {
		b.Cursor.DrawAndBack(" ")
		b.Cursor.MoveUp()
	}
	for b.Cursor.Position.X > 0 {
		b.Cursor.DrawAndBack(" ")
		b.Cursor.MoveLeft()
	}
	gu.RunCommand("[0m")
}

func (b *Board) CleanBoard() {
	gu.RunCommand("[2J")
}

func (b *Board) MoveSnake() {
	for i := 0; i < len(b.Snake.Body); i++ {
		oldPosition := b.Snake.Body[i]
		b.Cursor.MoveTo(oldPosition)
		b.Cursor.Draw(" ")
		var newPosition gu.Position
		switch b.Snake.Direction {
		case gu.UP:
			newPosition = gu.Position{X: oldPosition.X, Y: oldPosition.Y - 1}
		case gu.DOWN:
			newPosition = gu.Position{X: oldPosition.X, Y: oldPosition.Y + 1}
		case gu.LEFT:
			newPosition = gu.Position{X: oldPosition.X - 1, Y: oldPosition.Y}
		case gu.RIGHT:
			newPosition = gu.Position{X: oldPosition.X + 1, Y: oldPosition.Y}

		}
		b.Cursor.MoveTo(newPosition)
		b.Cursor.Draw(b.Snake.Pic)
	}
}
