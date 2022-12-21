package component

type Id struct {
	I int
}

func NewId(i int) Id {
	return Id{i}
}
