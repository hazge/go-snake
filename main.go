package main

import (
	"github.com/hazge/go-snake/board"
	"github.com/hazge/go-snake/gu"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	brd := board.New()
	brd.DrawBorder()
	for {
		brd.MoveSnake()
		time.Sleep(3 * time.Second)
	}

	defer func() {
		brd.Cursor.MoveTo(gu.Position{Y: brd.YLength})
	}()
}
