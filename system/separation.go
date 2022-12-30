package system

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Separation struct {
	*component.Boid
	*component.Facing
	*component.Id
	*component.Position
	*component.Radius
	*component.Target
	*component.TurnRate
}

func NewSeparation() *Separation {
	s := Separation{}

	return &s
}

func (s *Separation) Update(w engine.World) {
	separationRadius := 10.0

	var points = []num.Point[float64]{}

	//Move away from others - prevents stacking
	units := w.View(
		component.Position{},
		component.Radius{},
		component.Id{},
	).Filter()

	for _, u := range units {
		var pos *component.Position
		var rad *component.Radius
		var id *component.Id

		u.Get(&pos, &rad, &id)

		if id == s.Id {
			continue
		}

		if num.Distance(s.Position.X, s.Position.Y, pos.X, pos.Y) <= separationRadius {
			points = append(points, num.Point[float64]{X: pos.X, Y: pos.Y})
		}
	}

	//avg := num.AveragePoint(points...)

	//Move to target
	// if num.Distance(m.Position.X, m.Position.Y, m.Target.X, m.Target.Y) >= m.Speed.S {
	// 	if m.Position.X != m.Target.X && m.Position.Y != m.Target.Y {
	// 		mx := m.Target.X - m.Position.X
	// 		my := m.Target.Y - m.Position.Y
	// 		ang := math.Atan2(mx, my)

	// 		m.Position.X += math.Sin(ang) * m.Speed.S
	// 		m.Position.Y += math.Cos(ang) * m.Speed.S
	// 	}
	// }
}
