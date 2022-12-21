package entity

import "skharv/2DRTS/component"

type Unit struct {
	component.Color
	component.Id
	component.Position
	component.Radius
	component.Selected
	component.Speed
	component.Target
}
