package system

import (
	"math"
	"skharv/2DRTS/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Face struct {
	*component.Facing
	*component.Position
	*component.Target
	*component.TurnRate
}

func NewFace() *Face {
	f := Face{}

	return &f
}

func (f *Face) Update(w engine.World) {
	fx := f.Target.X - f.Position.X
	fy := f.Target.Y - f.Position.Y

	currentFacing := math.Mod((f.Facing.F + 360), 360)
	targetFacing := math.Atan2(fx, fy) * (180 / math.Pi)
	targetFacing = math.Mod((targetFacing + 360), 360)
	if targetFacing != currentFacing {
		netAngle := math.Mod((targetFacing - currentFacing + 360), 360)
		delta := math.Min(netAngle, f.TurnRate.R)
		sign := 0.0
		if netAngle < 180 {
			sign = 1
		} else {
			sign = -1
		}

		currentFacing += sign*delta + 360
		f.Facing.F = math.Mod((currentFacing), 360)
	}
}
