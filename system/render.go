package system

import (
	"image/color"
	"math"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/helper/num"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	offscreen *ebiten.Image
}

func NewRender() *Render {
	r := Render{}

	r.offscreen = ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight)

	return &r
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	//Color selected
	selection := w.View(
		component.Color{},
		component.Selected{},
	)

	selection.Each(func(e engine.Entity) {
		var col *component.Color
		var sel *component.Selected
		e.Get(&col, &sel)

		if sel.S {
			col.C = color.RGBA{255, 0, 0, 255}
		} else {
			col.C = color.RGBA{255, 255, 255, 255}
		}
	})

	//Draw Debug
	if globals.NavDebug {
		navmeshTris := w.View(
			component.Color{},
			component.NavMesh{},
			component.Triangle{},
		)

		navmeshTris.Each(func(e engine.Entity) {
			var col *component.Color
			var nav *component.NavMesh
			var tri *component.Triangle
			e.Get(&col, &nav, &tri)

			ebitenutil.DrawLine(r.offscreen, tri.A.X, tri.A.Y, tri.B.X, tri.B.Y, col.C)
			ebitenutil.DrawLine(r.offscreen, tri.B.X, tri.B.Y, tri.C.X, tri.C.Y, col.C)
			ebitenutil.DrawLine(r.offscreen, tri.C.X, tri.C.Y, tri.A.X, tri.A.Y, col.C)
		})

		navmeshPgon := w.View(
			component.Color{},
			component.NavMesh{},
			component.Polygon{},
		)

		navmeshPgon.Each(func(e engine.Entity) {
			var col *component.Color
			var nav *component.NavMesh
			var pol *component.Polygon
			e.Get(&col, &nav, &pol)

			for i, _ := range pol.P {
				if i < len(pol.P)-1 {
					ebitenutil.DrawLine(r.offscreen, pol.P[i].X, pol.P[i].Y, pol.P[i+1].X, pol.P[i+1].Y, col.C)
				} else {
					ebitenutil.DrawLine(r.offscreen, pol.P[i].X, pol.P[i].Y, pol.P[0].X, pol.P[0].Y, col.C)
				}
			}
		})
	}

	//Draw Units
	renders := w.View(
		component.Color{},
		component.Facing{},
		component.Position{},
		component.Radius{},
		component.Target{},
	)

	renders.Each(func(e engine.Entity) {
		var col *component.Color
		var fac *component.Facing
		var pos *component.Position
		var rad *component.Radius
		var tar *component.Target
		e.Get(&col, &fac, &pos, &rad, &tar)

		if globals.NavDebug {
			navmeshRects := w.View(
				component.NavMesh{},
				component.Triangle{},
			)

			navmeshRects.Each(func(e engine.Entity) {
				var nav *component.NavMesh
				var tri *component.Triangle
				e.Get(&nav, &tri)

				if num.PointInTriangle(num.Point[float64]{X: pos.X, Y: pos.Y}, tri.A, tri.B, tri.C) {
					col.C = color.RGBA{0, 255, 0, 128}
				}
			})
		}

		ebitenutil.DrawCircle(r.offscreen, pos.X, pos.Y, rad.R, col.C)

		if globals.Debug {
			c := color.RGBA{255, 255, 255, 255}
			//Line to Target
			ebitenutil.DrawLine(r.offscreen, pos.X, pos.Y, tar.X, tar.Y, c)

			//Line to Facing
			fx := 50 * math.Sin(fac.F*math.Pi/180)
			fy := 50 * math.Cos(fac.F*math.Pi/180)

			ebitenutil.DrawLine(r.offscreen, pos.X, pos.Y, pos.X+fx, pos.Y+fy, c)
		}
	})

	//Draw Cursor
	cursor := w.View(
		component.Clicked{},
		component.Color{},
		component.Rectangle{},
	)

	cursor.Each(func(e engine.Entity) {
		var cli *component.Clicked
		var col *component.Color
		var rec *component.Rectangle
		e.Get(&cli, &col, &rec)

		if cli.L {
			if rec.W != 0 || rec.H != 0 {
				ebitenutil.DrawRect(screen, rec.X, rec.Y, rec.W, rec.H, col.C)
			}
		}
	})

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()
}
