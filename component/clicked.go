package component

type Clicked struct {
	L bool
	R bool
	M bool
}

func NewClicked(L, R, M bool) Clicked {
	return Clicked{L, R, M}
}
