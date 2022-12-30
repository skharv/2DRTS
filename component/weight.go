package component

type Weight struct {
	W float64
}

func NewWeight(w float64) Weight {
	return Weight{w}
}
