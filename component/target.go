package component

type Target struct {
	X, Y float64
}

func NewTarget(x, y float64) Target {
	return Target{x, y}
}
