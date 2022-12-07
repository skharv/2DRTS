package scene

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/entity"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/system"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Clicked{},
		component.Color{},
		component.Position{},
		component.Radius{},
		component.Selected{},
	)

	w.AddSystems(
		system.NewCursor(),
		system.NewRender(),
		system.NewSelect(),
	)

	w.AddEntities(
		&entity.Unit{
			Color:    component.NewColor(255, 255, 255, 255),
			Position: component.NewPosition(globals.ScreenWidth/2, globals.ScreenHeight/2),
			Radius:   component.NewRadius(10),
			Selected: component.NewSelected(false),
		},
		&entity.Cursor{
			Clicked: component.NewClicked(false),
			Color:   component.NewColor(255, 255, 255, 255),
			Click:   component.NewPosition(0, 0),
			Current: component.NewPosition(0, 0),
		},
	)
}
