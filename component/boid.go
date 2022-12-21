package component

type Boid struct {
	L bool
	R bool
	M bool
}

func NewBoid(L, R, M bool) Boid {
	return Boid{L, R, M}
}
