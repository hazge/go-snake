package main

import (
	"github.com/hazge/snake/board"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	brd := board.New()
	brd.DrawBorder()
	time.Sleep(3 * time.Second)
	defer func() {
		brd.Cursor.MoveTo(brd.YLength, 0)
	}()
}
