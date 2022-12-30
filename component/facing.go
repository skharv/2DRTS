package component

type Facing struct {
	F float64
}

func NewFacing(f float64) Facing {
	return Facing{f}
}
