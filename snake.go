package main

import (
	"github.com/hazge/snake/board"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	b := board.New()
	b.DrawBorder()
	time.Sleep(3 * time.Second)
	defer func() {
		b.Cur.MoveTo(b.YLength, 0)
	}()
}
