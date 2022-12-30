package component

type Owner struct {
	O int
}

func NewOwner(o int) Owner {
	return Owner{o}
}
