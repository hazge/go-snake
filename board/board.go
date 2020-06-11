package board

import (
	"fmt"
	"github.com/hazge/snake/cursor"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	ESC   = "\033"
	CLEAR = "[2J"
)

type Board struct {
	Rows  uint16
	Lines uint16
	Cur   cursor.Cursor
}

func New() Board {
	w, h := getTerminalSize()
	//if w == 0 || h == 0 {
	//	log.Fatal("terminal empty")
	//}
	b := Board{w, h, cursor.New()}
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
	// left border
	for b.Cur.Position.Y < b.Rows {
		b.Cur.DrawAndBack("x")
		b.Cur.MoveDown()
	}
	for b.Cur.Position.X < b.Rows {
		b.Cur.DrawAndBack("x")
		b.Cur.MoveRight()
	}

}

func runCommand(command string) {
	fmt.Print(ESC + command)
}
func (b *Board) CleanBoard() {
	runCommand(CLEAR)
}
