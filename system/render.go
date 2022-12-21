package system

import (
	"image/color"
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"

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
	screen.Fill(color.Black)

	//Color selected
	selection := w.View(
		component.Selected{},
		component.Color{},
	).Filter()

	for _, e := range selection {
		var sel *component.Selected
		var col *component.Color
		e.Get(&sel, &col)

		if sel.S {
			col.C = color.RGBA{255, 0, 0, 255}
		} else {
			col.C = color.RGBA{255, 255, 255, 255}
		}
	}

	//Draw Units
	renders := w.View(
		component.Position{},
		component.Radius{},
		component.Color{},
	).Filter()

	for _, e := range renders {
		var pos *component.Position
		var rad *component.Radius
		var col *component.Color
		e.Get(&pos, &rad, &col)

		ebitenutil.DrawCircle(r.offscreen, pos.X, pos.Y, rad.R, col.C)
	}

	//Draw Cursor
	cursor := w.View(
		component.Clicked{},
		component.Color{},
		component.Rectangle{},
	).Filter()

	for _, e := range cursor {
		var cli *component.Clicked
		var col *component.Color
		var rec *component.Rectangle
		e.Get(&cli, &col, &rec)

		if cli.L {
			if rec.W != 0 || rec.H != 0 {
				ebitenutil.DrawRect(screen, rec.X, rec.Y, rec.W, rec.H, col.C)
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()
}
