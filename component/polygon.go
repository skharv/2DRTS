package component

import "skharv/2DRTS/helper/num"

type Polygon struct {
	P []num.Point[float64]
}

func NewPolygon(p ...num.Point[float64]) Polygon {
	return Polygon{p}
}
