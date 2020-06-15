package snake

import (
	"github.com/hazge/go-snake/gu"
	"math/rand"
)

type Snake struct {
	Pic       string
	Body      []gu.Position
	Direction string
}

func New(x, y uint16) Snake {
	x = uint16(rand.Intn(int(x)))
	y = uint16(rand.Intn(int(y)))
	return Snake{Body: []gu.Position{{X: x, Y: y}}, Direction: gu.RIGHT, Pic: "·"}
}
