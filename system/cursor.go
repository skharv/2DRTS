package system

import (
	"skharv/2DRTS/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Cursor struct {
	*component.Color
	*component.Clicked
	Click   *component.Position
	Current *component.Position
}

func NewCursor() *Cursor {
	c := Cursor{}

	return &c
}

func (c *Cursor) Update(w engine.World) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !c.Clicked.C {
			c.Clicked.C = true
			cx, cy := ebiten.CursorPosition()
			c.Click.X = float64(cx)
			c.Click.Y = float64(cy)
		}
	}

	cx, cy := ebiten.CursorPosition()
	c.Current.X = float64(cx)
	c.Current.Y = float64(cy)
}
