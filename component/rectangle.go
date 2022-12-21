package component

type Rectangle struct {
	X, Y, W, H float64
}

func NewRectangle(x, y, w, h float64) Rectangle {
	return Rectangle{x, y, w, h}
}
