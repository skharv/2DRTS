package manager

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"

	"github.com/sedyh/mizu/pkg/engine"
)

type Chunk struct {
	Entities []engine.Entity
}

var (
	Chunks = [][]Chunk{}
)

func InitialiseChunks(w engine.World) {
	maxX := globals.ScreenWidth / globals.ChunkSize
	maxY := globals.ScreenHeight / globals.ChunkSize

	var chunks = [][]Chunk{}

	for i := 0; i < maxX; i++ {
		c := []Chunk{}
		for j := 0; j < maxY; j++ {
			ch := Chunk{}
			e := []engine.Entity{}
			ch.Entities = e
			c = append(c, ch)
		}
		chunks = append(chunks, c)
	}

	units := w.View(
		component.Chunk{},
		component.Id{},
		component.Position{},
	)

	units.Each(func(e engine.Entity) {
		var chu *component.Chunk
		var id *component.Id
		var pos *component.Position

		e.Get(&chu, &id, &pos)

		cPosX := int(pos.X / globals.ChunkSize)
		cPosY := int(pos.Y / globals.ChunkSize)

		chu.X = cPosX
		chu.Y = cPosY
		chunks[cPosX][cPosY].Entities = append(chunks[cPosX][cPosY].Entities, e)
	})

	Chunks = chunks
}

func (c *Chunk) MoveUnit(entityId int, newChunk *Chunk) {
	for i, e := range c.Entities {
		var id *component.Id
		e.Get(&id)

		if id.I != entityId {
			continue
		}

		newChunk.Entities = append(newChunk.Entities, e)

		c.Entities[i] = c.Entities[len(c.Entities)-1]
		c.Entities[len(c.Entities)-1] = nil
		c.Entities = c.Entities[:len(c.Entities)-1]
		return
	}
}
