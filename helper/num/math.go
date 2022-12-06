package num

import "math"

func Distance(p1x, p1y, p2x, p2y float64) float64 {
	d := (p2x*p2x - p1x*p1x) + (p2y*p2y - p1y*p1y)
	return math.Sqrt(d)
}
