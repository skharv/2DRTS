package system

import (
	"math"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/helper/num"

	"github.com/sedyh/mizu/pkg/engine"
)

type Accelerate struct {
	*component.Facing
	*component.Id
	*component.Position
	*component.Radius
	*component.Speed
	*component.State
	*component.Target
	*component.Velocity
}

func NewAccelerate() *Accelerate {
	a := Accelerate{}

	return &a
}

func (a *Accelerate) Update(w engine.World) {
	if a.State.S == globals.Move {
		if num.Distance(a.Position.X, a.Position.Y, a.Target.X, a.Target.Y) > a.Radius.R {
			if a.Position.X != a.Target.X || a.Position.Y != a.Target.Y {
				posX := math.Sin(a.Facing.F*math.Pi/180) * a.Speed.S
				posY := math.Cos(a.Facing.F*math.Pi/180) * a.Speed.S

				a.Velocity.X += posX
				a.Velocity.Y += posY
			}
		} else {
			a.State.S = globals.Idle
		}
	}
}
