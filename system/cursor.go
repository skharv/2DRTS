package system

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/num"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Cursor struct {
	*component.Color
	*component.Clicked
	*component.Rectangle
}

func NewCursor() *Cursor {
	c := Cursor{}

	return &c
}

func (c *Cursor) Update(w engine.World) {
	//Left button
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !c.Clicked.L {
			c.Clicked.L = true
			cx, cy := ebiten.CursorPosition()
			c.Rectangle.X = float64(cx)
			c.Rectangle.Y = float64(cy)
		}
	}

	if c.Clicked.L {
		cx, cy := ebiten.CursorPosition()
		c.Rectangle.W = float64(cx) - c.Rectangle.X
		c.Rectangle.H = float64(cy) - c.Rectangle.Y
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		c.Clicked.L = false
	}

	units := w.View(
		component.Position{},
		component.Radius{},
		component.Color{},
		component.Selected{},
		component.Target{},
	).Filter()

	for _, e := range units {
		var pos *component.Position
		var rad *component.Radius
		var col *component.Color
		var sel *component.Selected
		var tar *component.Target

		e.Get(&pos, &rad, &col, &sel, &tar)

		if c.Clicked.L {
			if c.Rectangle.W != 0 || c.Rectangle.H != 0 {
				if num.CircleRectangleIntersect(
					pos.X,
					pos.Y,
					rad.R,
					c.Rectangle.X,
					c.Rectangle.Y,
					c.Rectangle.W,
					c.Rectangle.H) {
					sel.S = true
				} else {
					sel.S = false
				}
			} else {
				if num.Distance(pos.X, pos.Y, c.Rectangle.X, c.Rectangle.Y) < rad.R {
					sel.S = true
				} else {
					sel.S = false
				}
			}
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && sel.S {
			rx, ry := ebiten.CursorPosition()
			tar.X = float64(rx)
			tar.Y = float64(ry)
		}
	}
}
