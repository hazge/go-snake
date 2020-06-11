package board

import (
	"github.com/hazge/snake/cursor"
	"github.com/hazge/snake/gu"
	"github.com/hazge/snake/snake"
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
}

func New() Board {
	w, h := getTerminalSize()
	b := Board{w, h, cursor.New(), snake.New(w, h)}
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
	for b.Cur.Position.Y < b.YLength {
		b.Cur.DrawAndBack(" ")
		b.Cur.MoveDown()
	}

	for b.Cur.Position.X < b.XLength {
		b.Cur.DrawAndBack(" ")
		b.Cur.MoveRight()
	}
	for b.Cur.Position.Y > 0 {
		b.Cur.DrawAndBack(" ")
		b.Cur.MoveUp()
	}
	for b.Cur.Position.X > 0 {
		b.Cur.DrawAndBack(" ")
		b.Cur.MoveLeft()
	}
	gu.RunCommand("[0m")
}

func (b *Board) CleanBoard() {
	gu.RunCommand("[2J")
}
