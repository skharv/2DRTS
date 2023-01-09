package entity

import "skharv/2DRTS/component"

type Building struct {
	component.Chunk
	component.Color
	component.Id
	component.Owner
	component.Position
	component.Radius
	component.Selected
	component.State
	component.Target
}
