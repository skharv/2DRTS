package component

import "skharv/2DRTS/helper/globals"

type State struct {
	S globals.State
}

func NewState(s globals.State) State {
	return State{s}
}
