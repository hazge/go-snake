package cursor

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

type Position struct {
	X uint16
	Y uint16
}
type Cursor struct {
	Position Position
}

func New() Cursor {
	c := Cursor{Position: Position{0, 0}}
	c.MoveTo(0, 0)
	return c
}
func (c *Cursor) Draw(text string) {
	fmt.Print(text)
	c.Position.X += uint16(len(text))
}

func (c *Cursor) DrawAndBack(text string) {
	c.Draw(text)
	c.MoveLeft()
}

func runCommand(command string) {
	fmt.Print(ESC + command)
}

func (c *Cursor) MoveTo(x, y uint16) {
	runCommand(fmt.Sprintf("[%d;%df", x, y))
}

func (c *Cursor) MoveUp() {
	runCommand(fmt.Sprintf("[A"))
	c.Position.Y -= 1

}

func (c *Cursor) MoveDown() {
	runCommand(fmt.Sprintf("[B"))
	c.Position.Y += 1
}

func (c *Cursor) MoveLeft() {
	runCommand(fmt.Sprintf("[D"))
	c.Position.X -= 1

}

func (c *Cursor) MoveRight() {
	runCommand(fmt.Sprintf("[C"))
	c.Position.X += 1
}
