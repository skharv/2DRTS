package globals

const (
	ScreenWidth  = 1024
	ScreenHeight = 768
	ChunkSize    = 32
	Debug        = false
	NavDebug     = true
	P1Owner      = 1
)

type State int64

const (
	Idle State = iota
	Move
	Attack
	Follow
)
