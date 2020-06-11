package cursor

import (
	"fmt"
	"github.com/hazge/snake/gu"
)

type Cursor struct {
	Position gu.Position
}

func New() Cursor {
	c := Cursor{Position: gu.Position{}}
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

func (c *Cursor) MoveTo(x, y uint16) {
	gu.RunCommand(fmt.Sprintf("[%d;%df", x, y))
}

func (c *Cursor) MoveUp() {
	gu.RunCommand("[A")
	c.Position.Y -= 1

}

func (c *Cursor) MoveDown() {
	gu.RunCommand("[B")
	c.Position.Y += 1
}

func (c *Cursor) MoveLeft() {
	gu.RunCommand("[D")
	c.Position.X -= 1

}

func (c *Cursor) MoveRight() {
	gu.RunCommand("[C")
	c.Position.X += 1
}
