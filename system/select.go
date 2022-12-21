package system

import (
	"skharv/2DRTS/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Select struct {
	*component.Selected
	*component.Position
	*component.Radius
}

func NewSelect() *Select {
	s := Select{}

	return &s
}

func (s *Select) Update(w engine.World) {
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	cx, cy := ebiten.CursorPosition()
	// 	if num.Distance(s.Position.X, s.Position.Y, float64(cx), float64(cy)) < s.Radius.R {
	// 		s.Selected.S = true
	// 	} else {
	// 		s.Selected.S = false
	// 	}
	// }
}
