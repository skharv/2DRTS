package num

import "math"

func Distance(p1x, p1y, p2x, p2y float64) float64 {
	d := math.Pow(p2x-p1x, 2) + math.Pow(p2y-p1y, 2)
	return math.Sqrt(d)
}

func CircleRectangleIntersect(cx, cy, cr, rx, ry, rw, rh float64) bool {
	if PointInRectangle(cx, cy, rx, ry, rw, rh) {
		return true
	}
	if LineInCircle(cx, cy, cr, rx, ry, rx+rw, ry) {
		return true
	}
	if LineInCircle(cx, cy, cr, rx, ry, rx, ry+rh) {
		return true
	}
	if LineInCircle(cx, cy, cr, rx, ry+rh, rx+rw, ry+rh) {
		return true
	}
	if LineInCircle(cx, cy, cr, rx+rw, ry, rx+rw, ry+rh) {
		return true
	}

	return false
}

func InfiniteLineThroughCircle(cx, cy, cr, p1x, p1y, p2x, p2y float64) bool {

	x := (math.Abs(((p2x-p1x)*cy)+(cx*(p1y-p2y))+
		((p1x-p2x)*p1y)+((p1y-p2y)*p1x))/
		math.Sqrt(math.Pow(p2x-p1x, 2)+math.Pow(p1y-p2y, 2)) <= cr)

	return x
}

func LineInCircle(cx, cy, cr, p1x, p1y, p2x, p2y float64) bool {
	p1x -= cx
	p1y -= cy
	p2x -= cx
	p2y -= cy

	a := math.Pow(p2x-p1x, 2) + math.Pow(p2y-p1y, 2)
	b := 2 * (p1x*(p2x-p1x) + p1y*(p2y-p1y))
	c := math.Pow(p1x, 2) + math.Pow(p1y, 2) - math.Pow(cr, 2)
	disc := math.Pow(b, 2) - 4*a*c
	if disc <= 0 {
		return false
	}

	sqrtdisc := math.Sqrt(disc)

	t1 := (-b + sqrtdisc) / (2 * a)
	t2 := (-b - sqrtdisc) / (2 * a)

	if (0 < t1 && t1 < 1) || (0 < t2 && t2 < 1) {
		return true
	}

	return false
}

func PointInRectangle(px, py, rx, ry, rw, rh float64) bool {
	//needs abs
	minX := rx
	maxX := rx + rw

	if maxX < minX {
		minX = maxX
		maxX = rx
	}

	minY := ry
	maxY := ry + rh

	if maxY < minY {
		minY = maxY
		maxY = ry
	}

	if px > minX && px <= maxX {
		if py > minY && py <= maxY {
			return true
		}
	}
	return false
}
