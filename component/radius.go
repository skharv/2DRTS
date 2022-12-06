package component

type Radius struct {
	R float64
}

func NewRadius(r float64) Radius {
	return Radius{r}
}
