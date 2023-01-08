package system

import (
	"skharv/2DRTS/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Move struct {
	*component.Position
	*component.Velocity
}

func NewMove() *Move {
	m := Move{}

	return &m
}

func (m *Move) Update(w engine.World) {
	if m.Velocity.X != 0.0 || m.Velocity.Y != 0.0 {
		m.Position.X += m.Velocity.X
		m.Position.Y += m.Velocity.Y

		m.Velocity.X = 0.0
		m.Velocity.Y = 0.0
	}
}
