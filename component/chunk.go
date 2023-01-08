package component

type Chunk struct {
	X, Y int
}

func NewChunk(x, y int) Chunk {
	return Chunk{x, y}
}
