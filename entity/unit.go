package entity

import "skharv/2DRTS/component"

type Unit struct {
	component.Color
	component.Facing
	component.Id
	component.Owner
	component.Position
	component.Radius
	component.Selected
	component.Speed
	component.State
	component.Target
	component.TurnRate
	component.Weight
}
