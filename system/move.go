package system

import (
	"math"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Move struct {
	*component.Facing
	*component.Id
	*component.Position
	*component.Radius
	*component.Speed
	*component.State
	*component.Target
	*component.Weight
}

func NewMove() *Move {
	m := Move{}

	return &m
}

func (m *Move) Update(w engine.World) {
	//Move away from others - prevents stacking

	units := w.View(
		component.Id{},
		component.Position{},
		component.Radius{},
		component.Weight{},
	).Filter()

	for _, u := range units {
		var id *component.Id
		var pos *component.Position
		var rad *component.Radius
		var wei *component.Weight

		u.Get(&id, &pos, &rad, &wei)

		if id == m.Id {
			continue
		}

		if num.Distance(m.Position.X, m.Position.Y, pos.X, pos.Y) <= m.Radius.R+rad.R {
			mx := pos.X - m.Position.X
			my := pos.Y - m.Position.Y
			ang := math.Atan2(mx, my) + math.Pi
			strength := m.Weight.W + wei.W
			strength = wei.W / strength

			m.Position.X += math.Sin(ang) * m.Speed.S * strength
			m.Position.Y += math.Cos(ang) * m.Speed.S * strength
		}
	}

	//Move to target
	if m.State.S == globals.Move {
		if num.Distance(m.Position.X, m.Position.Y, m.Target.X, m.Target.Y) > m.Radius.R {
			if m.Position.X != m.Target.X && m.Position.Y != m.Target.Y {
				m.Position.X += math.Sin(m.Facing.F*math.Pi/180) * m.Speed.S
				m.Position.Y += math.Cos(m.Facing.F*math.Pi/180) * m.Speed.S
			}
		} else {
			m.State.S = globals.Idle
		}
	}
}
