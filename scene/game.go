package scene

import (
	"math/rand"
	"skharv/2DRTS/component"
	"skharv/2DRTS/entity"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/prefabs"
	"skharv/2DRTS/system"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Clicked{},
		component.Color{},
		component.Facing{},
		component.Id{},
		component.NavMesh{},
		component.Owner{},
		component.Position{},
		component.Radius{},
		component.Rectangle{},
		component.Selected{},
		component.Speed{},
		component.State{},
		component.Target{},
		component.TurnRate{},
		component.Weight{},
	)

	w.AddSystems(
		system.NewCursor(),
		system.NewFace(),
		system.NewMove(),
		system.NewRender(),
		system.NewSelect(),
	)

	w.AddEntities(
		&entity.Cursor{
			Clicked:   component.NewClicked(false, false, false),
			Color:     component.NewColor(0, 255, 0, 128),
			Rectangle: component.NewRectangle(0, 0, 0, 0),
		},
	)

	entities := prefabs.Initialise()

	for i := 0; i < 100; i++ {
		x := float64(rand.Intn(globals.ScreenWidth))
		y := float64(rand.Intn(globals.ScreenHeight))

		if rand.Intn(2) < 1 {
			w.AddEntities(entities.SpawnUnit("unitA", x, y, rand.Intn(2)))
		} else {
			w.AddEntities(entities.SpawnUnit("unitB", x, y, rand.Intn(2)))
		}
	}
}
