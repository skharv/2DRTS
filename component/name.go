package component

type Name struct {
	N string
}

func NewName(n string) Name {
	return Name{n}
}
