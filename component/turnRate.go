package component

type TurnRate struct {
	R float64
}

func NewTurnRate(r float64) TurnRate {
	return TurnRate{r}
}
