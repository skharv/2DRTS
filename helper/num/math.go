package num

import "math"

func Distance(p1x, p1y, p2x, p2y float64) float64 {
	d := math.Pow(p2x-p1x, 2) + math.Pow(p2y-p1y, 2)
	return math.Sqrt(d)
}
