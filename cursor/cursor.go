package cursor

import (
	"fmt"
	"github.com/hazge/go-snake/gu"
)

type Cursor struct {
	Position gu.Position
}

func New() Cursor {
	c := Cursor{Position: gu.Position{}}
	c.MoveTo(gu.Position{})
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

func (c *Cursor) MoveTo(p gu.Position) {
	gu.RunCommand(fmt.Sprintf("[%d;%df", p.X, p.Y))
}

func (c *Cursor) MoveUp() {
	gu.RunCommand(gu.UP)
	c.Position.Y -= 1

}

func (c *Cursor) MoveDown() {
	gu.RunCommand(gu.DOWN)
	c.Position.Y += 1
}

func (c *Cursor) MoveLeft() {
	gu.RunCommand(gu.LEFT)
	c.Position.X -= 1

}

func (c *Cursor) MoveRight() {
	gu.RunCommand(gu.RIGHT)
	c.Position.X += 1
}
