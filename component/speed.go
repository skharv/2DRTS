package component

type Speed struct {
	S float64
}

func NewSpeed(s float64) Speed {
	return Speed{s}
}
