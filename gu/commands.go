package gu

import "fmt"

const (
	ESC           = "\033"
	MOVE_TOP_LEFT = "[H"
	CLEAR         = "[2J"
)
const (
	tN = iota
	tW
	tS
	tE
)

func RunCommand(command string) {
	fmt.Print(ESC + command)
}
