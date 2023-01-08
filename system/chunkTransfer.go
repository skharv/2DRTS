package system

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/helper/globals"
	"skharv/2DRTS/helper/manager"

	"github.com/sedyh/mizu/pkg/engine"
)

type ChunkTransfer struct {
	*component.Chunk
	*component.Id
	*component.Position
}

func NewChunkTransfer() *ChunkTransfer {
	c := ChunkTransfer{}

	return &c
}

func (c *ChunkTransfer) Update(w engine.World) {
	cPosX := int(c.Position.X / globals.ChunkSize)
	cPosY := int(c.Position.Y / globals.ChunkSize)

	if cPosX < 0 || cPosX >= len(manager.Chunks) {
		return
	}
	if cPosY < 0 || cPosY >= len(manager.Chunks[cPosX]) {
		return
	}
	if cPosX != c.Chunk.X || cPosY != c.Chunk.Y {
		manager.Chunks[c.Chunk.X][c.Chunk.Y].MoveUnit(c.Id.I, &manager.Chunks[cPosX][cPosY])

		c.Chunk.X = cPosX
		c.Chunk.Y = cPosY
	}
}
