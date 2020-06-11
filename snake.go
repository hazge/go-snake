package main

import (
	"github.com/hazge/snake/board"
	"time"
)

func main() {
	b := board.New()
	b.DrawBorder()
	time.Sleep(10 * time.Second)
}
