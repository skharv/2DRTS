package system

import (
	"math"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Move struct {
	*component.Speed
	*component.Position
	*component.Target
}

func NewMove() *Move {
	m := Move{}

	return &m
}

func (m *Move) Update(w engine.World) {
	if num.Distance(m.Position.X, m.Position.Y, m.Target.X, m.Target.Y) >= m.Speed.S {
		if m.Position.X != m.Target.X && m.Position.Y != m.Target.Y {
			mx := m.Target.X - m.Position.X
			my := m.Target.Y - m.Position.Y
			ang := math.Atan2(mx, my)

			m.Position.X += math.Sin(ang) * m.Speed.S
			m.Position.Y += math.Cos(ang) * m.Speed.S
		}
	}
}
