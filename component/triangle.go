package component

import "skharv/2DRTS/helper/num"

type Triangle struct {
	A, B, C num.Point[float64]
}

func NewTriangle(a, b, c num.Point[float64]) Triangle {
	return Triangle{a, b, c}
}
