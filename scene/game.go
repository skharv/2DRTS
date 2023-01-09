package scene

import (
	"math/rand"
	"skharv/2DRTS/component"
	"skharv/2DRTS/entity"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/helper/manager"
	"skharv/2DRTS/helper/num"
	"skharv/2DRTS/prefabs"
	"skharv/2DRTS/system"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Chunk{},
		component.Clicked{},
		component.Color{},
		component.Facing{},
		component.Id{},
		component.Name{},
		component.NavMesh{},
		component.Owner{},
		component.Position{},
		component.Radius{},
		component.Rectangle{},
		component.Selected{},
		component.Speed{},
		component.State{},
		component.Target{},
		component.Triangle{},
		component.TurnRate{},
		component.Velocity{},
		component.Weight{},
	)

	w.AddSystems(
		system.NewAccelerate(),
		system.NewChunkTransfer(),
		system.NewCursor(),
		system.NewFace(),
		system.NewMove(),
		system.NewRender(),
		system.NewShove(),
	)

	w.AddEntities(
		&entity.Cursor{
			Clicked:   component.NewClicked(false, false, false),
			Color:     component.NewColor(0, 255, 0, 128),
			Rectangle: component.NewRectangle(0, 0, 0, 0),
		},
		&entity.WalkableSpace{
			Color:   component.NewColor(0, 255, 0, 128),
			NavMesh: component.NewNavMesh(),
			Triangle: component.NewTriangle(
				num.Point[float64]{X: 50, Y: 50},
				num.Point[float64]{X: 200, Y: 100},
				num.Point[float64]{X: 100, Y: 300},
			),
		},
		&entity.WalkableSpace{
			Color:   component.NewColor(0, 255, 0, 128),
			NavMesh: component.NewNavMesh(),
			Triangle: component.NewTriangle(
				num.Point[float64]{X: 200, Y: 100},
				num.Point[float64]{X: 100, Y: 300},
				num.Point[float64]{X: 400, Y: 500},
			),
		},
		&entity.WalkableSpace{
			Color:   component.NewColor(0, 255, 0, 128),
			NavMesh: component.NewNavMesh(),
			Triangle: component.NewTriangle(
				num.Point[float64]{X: 200, Y: 100},
				num.Point[float64]{X: 400, Y: 500},
				num.Point[float64]{X: 700, Y: 300},
			),
		},
		&entity.WalkableSpace{
			Color:   component.NewColor(0, 255, 0, 128),
			NavMesh: component.NewNavMesh(),
			Triangle: component.NewTriangle(
				num.Point[float64]{X: 400, Y: 500},
				num.Point[float64]{X: 700, Y: 300},
				num.Point[float64]{X: 800, Y: 600},
			),
		},
	)

	entities := prefabs.Entities{}

	for i := 0; i < 400; i++ {
		x := rand.Intn(globals.ScreenWidth)
		y := rand.Intn(globals.ScreenHeight)

		// x = 100
		// y = 100

		switch rand.Intn(4) {
		case 0:
			w.AddEntities(entities.SpawnUnit("unitA", x, y, 1))
		case 1:
			w.AddEntities(entities.SpawnUnit("unitB", x, y, 1))
		case 2:
			w.AddEntities(entities.SpawnUnit("unitC", x, y, 1))
		case 3:
			w.AddEntities(entities.SpawnUnit("unitD", x, y, 1))
		}
	}

	manager.InitialiseChunks(w)
}
