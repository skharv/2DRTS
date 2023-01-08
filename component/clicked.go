package component

type Clicked struct {
	L bool
	R bool
	M bool
}

func NewClicked(l, r, m bool) Clicked {
	return Clicked{l, r, m}
}
