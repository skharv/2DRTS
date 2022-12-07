package component

type Clicked struct {
	C bool
}

func NewClicked(c bool) Clicked {
	return Clicked{c}
}
