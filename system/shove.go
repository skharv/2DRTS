package system

import (
	"math"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/manager"
	"skharv/2DRTS/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Shove struct {
	*component.Chunk
	*component.Facing
	*component.Id
	*component.Position
	*component.Radius
	*component.Speed
	*component.State
	*component.Selected
	*component.Target
	*component.Velocity
	*component.Weight
}

func NewShove() *Shove {
	s := Shove{}

	return &s
}

func (m *Shove) Update(w engine.World) {
	var nearby = []engine.Entity{}

	for i := -1; i < 1; i++ {
		x := m.Chunk.X + i
		if x < 0 || x > len(manager.Chunks) {
			continue
		}
		for j := -1; j < 1; j++ {
			y := m.Chunk.Y + j
			if y < 0 || y > len(manager.Chunks[x]) {
				continue
			}

			nearby = append(nearby, manager.Chunks[x][y].Entities...)
		}
	}

	for _, e := range nearby {
		var id *component.Id
		var pos *component.Position
		var rad *component.Radius
		var spe *component.Speed
		var vel *component.Velocity
		var wei *component.Weight

		e.Get(&id, &pos, &rad, &spe, &vel, &wei)

		if id.I == m.Id.I {
			continue
		}

		dist := num.Distance(m.Position.X, m.Position.Y, pos.X, pos.Y)
		combinedSize := m.Radius.R + rad.R

		if dist <= combinedSize {
			mx := pos.X - m.Position.X
			my := pos.Y - m.Position.Y
			ang := math.Atan2(mx, my) + math.Pi
			total := m.Weight.W + wei.W
			self := wei.W / total
			other := m.Weight.W / total

			m.Velocity.X += math.Sin(ang) * m.Speed.S * self
			m.Velocity.Y += math.Cos(ang) * m.Speed.S * self

			vel.X += math.Sin(ang+math.Pi) * m.Speed.S * other
			vel.Y += math.Cos(ang+math.Pi) * m.Speed.S * other
		}
	}
}
