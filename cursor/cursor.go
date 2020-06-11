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
	X int
	Y int
}
type Cursor struct {
	position Position
}

func New() Cursor {
	c := Cursor{position: Position{1, 1}}
	c.MoveTo(0, 0)
	return c
}
func (c *Cursor) Draw(text string) {
	fmt.Print(text)
	c.position.Y += len(text)
}

func (c *Cursor) DrawAndBack(text string) {
	fmt.Print(text)
	c.MoveLeft()
}

func runCommand(command string) {
	fmt.Print(ESC + command)
}

func (c *Cursor) MoveTo(row, column int) {
	runCommand(fmt.Sprintf("[%d;%df", row, column))
}

func (c *Cursor) MoveUp() {
	runCommand(fmt.Sprintf("[A"))
	c.position.X -= 1

}

func (c *Cursor) MoveDown() {
	runCommand(fmt.Sprintf("[B"))
	c.position.X += 1
}

func (c *Cursor) MoveLeft() {
	runCommand(fmt.Sprintf("[D"))
	c.position.Y -= 1

}

func (c *Cursor) MoveRight() {
	runCommand(fmt.Sprintf("[C"))
	c.position.Y += 1
}
