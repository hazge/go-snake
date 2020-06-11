package board

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"maze/cursor"
	"os"
)

const (
	ESC   = "\033"
	CLEAR = "[2J"
)

type Board struct {
	Rows  int
	Lines int
	C     cursor.Cursor
}

func New() Board {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic("can't access terminal.")
	}
	b := Board{w, h, cursor.New()}
	b.CleanBoard()
	return b

}
func (b *Board) DrawBorder() {
	for y := 1; y <= b.Rows; y++ {
		b.C.Draw("y")
	}
}

func runCommand(command string) {
	fmt.Print(ESC + command)
}
func (b *Board) CleanBoard() {
	runCommand(CLEAR)
}
