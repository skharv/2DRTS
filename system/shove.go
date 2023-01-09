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

func (s *Shove) Update(w engine.World) {
	var nearby = []engine.Entity{}

	for i := -1; i < 1; i++ {
		x := s.Chunk.X + i
		if x < 0 || x > len(manager.Chunks) {
			continue
		}
		for j := -1; j < 1; j++ {
			y := s.Chunk.Y + j
			if y < 0 || y > len(manager.Chunks[x]) {
				continue
			}

			nearby = append(nearby, manager.Chunks[x][y].Entities...)
		}
	}

	s.Bump(nearby)
}

func (s *Shove) Bump(nearby []engine.Entity) {
	for _, e := range nearby {
		var id *component.Id
		var pos *component.Position
		var rad *component.Radius
		var spe *component.Speed
		var vel *component.Velocity
		var wei *component.Weight

		e.Get(&id, &pos, &rad, &spe, &vel, &wei)

		if id.I == s.Id.I {
			continue
		}

		dist := num.Distance(s.Position.X, s.Position.Y, pos.X, pos.Y)
		combinedSize := s.Radius.R + rad.R

		if dist <= combinedSize {
			mx := pos.X - s.Position.X
			my := pos.Y - s.Position.Y
			ang := math.Atan2(mx, my) + math.Pi
			total := s.Weight.W + wei.W
			self := wei.W / total
			other := s.Weight.W / total

			s.Velocity.X += math.Sin(ang) * s.Speed.S * self
			s.Velocity.Y += math.Cos(ang) * s.Speed.S * self

			vel.X += math.Sin(ang+math.Pi) * s.Speed.S * other
			vel.Y += math.Cos(ang+math.Pi) * s.Speed.S * other
		}
	}
}

func (s *Shove) Separation(nearby []engine.Entity) {

}

func (s *Shove) Alignment(nearby []engine.Entity) {

}

func (s *Shove) Cohesion(nearby []engine.Entity) {

}
