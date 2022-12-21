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
		component.Rectangle{},
		component.Selected{},
		component.Speed{},
		component.Target{},
	)

	w.AddSystems(
		system.NewCursor(),
		system.NewMove(),
		system.NewRender(),
		system.NewSelect(),
	)

	w.AddEntities(
		&entity.Unit{
			Color:    component.NewColor(255, 255, 255, 255),
			Position: component.NewPosition(globals.ScreenWidth/2, globals.ScreenHeight/2),
			Radius:   component.NewRadius(10),
			Selected: component.NewSelected(false),
			Speed:    component.NewSpeed(2),
			Target:   component.NewTarget(globals.ScreenWidth/2, globals.ScreenHeight/2),
		},
		&entity.Cursor{
			Clicked:   component.NewClicked(false, false, false),
			Color:     component.NewColor(0, 255, 0, 128),
			Rectangle: component.NewRectangle(0, 0, 0, 0),
		},
	)
}
