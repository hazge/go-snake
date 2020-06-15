package gu

import "fmt"

const (
	ESC           = "\033"
	MOVE_TOP_LEFT = "[H"
	CLEAR         = "[2J"
)

func RunCommand(command string) {
	fmt.Print(ESC + command)
}

const (
	UP    = "[A"
	DOWN  = "[B"
	RIGHT = "[C"
	LEFT  = "[D"
)
