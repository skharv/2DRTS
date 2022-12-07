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
		component.Color{},
		component.Position{},
		component.Radius{},
		component.Selected{},
	)

	w.AddSystems(
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
		&entity.Unit{
			Color:    component.NewColor(255, 255, 255, 255),
			Position: component.NewPosition((globals.ScreenWidth/4)*3, (globals.ScreenHeight/4)*3),
			Radius:   component.NewRadius(10),
			Selected: component.NewSelected(false),
		},
		&entity.Unit{
			Color:    component.NewColor(255, 255, 255, 255),
			Position: component.NewPosition(globals.ScreenWidth/4, globals.ScreenHeight/4),
			Radius:   component.NewRadius(10),
			Selected: component.NewSelected(false),
		},
	)
}
