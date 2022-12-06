package component

type Selected struct {
	S bool
}

func NewSelected(s bool) Selected {
	return Selected{s}
}
