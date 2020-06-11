package snake

import (
	"github.com/hazge/snake/gu"
	"math/rand"
)

type Snake struct {
	Len  int
	Body []gu.Position
}

func New(x, y uint16) Snake {
	x = uint16(rand.Intn(int(x)))
	y = uint16(rand.Intn(int(y)))
	return Snake{1, []gu.Position{{X: x, Y: y}}}
}
