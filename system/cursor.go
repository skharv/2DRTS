package system

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"
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

	// var ctrlnam *component.Name
	// var ctrlown *component.Owner

	units := w.View(
		component.Color{},
		component.Name{},
		component.Owner{},
		component.Position{},
		component.Radius{},
		component.Selected{},
		component.State{},
		component.Target{},
	)

	units.Each(func(e engine.Entity) {
		var col *component.Color
		var nam *component.Name
		var own *component.Owner
		var pos *component.Position
		var rad *component.Radius
		var sel *component.Selected
		var sta *component.State
		var tar *component.Target

		e.Get(&col, &nam, &own, &pos, &rad, &sel, &sta, &tar)

		if own.O != globals.P1Owner {
			return
		}

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
					if !ebiten.IsKeyPressed(ebiten.KeyShift) {
						sel.S = false
					}
				}
			} else {
				if num.Distance(pos.X, pos.Y, c.Rectangle.X, c.Rectangle.Y) < rad.R {
					// if ebiten.IsKeyPressed(ebiten.KeyControl) {
					// 	ctrlnam = nam
					// 	ctrlown = own
					// }
					sel.S = true
				} else {
					if !ebiten.IsKeyPressed(ebiten.KeyShift) {
						sel.S = false
					}
				}
			}
		}

		// if ctrlnam != nil && ctrlown != nil {
		// 	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		// 		similar := w.View(
		// 			component.Name{},
		// 			component.Owner{},
		// 			component.Selected{},
		// 		)

		// 		similar.Each(func(e engine.Entity) {
		// 			var nam *component.Name
		// 			var own *component.Owner
		// 			var sel *component.Selected

		// 			e.Get(&nam, &own, &sel)

		// 			if own.O != ctrlown.O {
		// 				return
		// 			}
		// 			if nam.N != ctrlnam.N {
		// 				return
		// 			}

		// 			sel.S = true
		// 		})
		// 	}
		// }

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) && sel.S {
			rx, ry := ebiten.CursorPosition()
			tar.X = float64(rx)
			tar.Y = float64(ry)
			sta.S = globals.Move
		}
	})
}
