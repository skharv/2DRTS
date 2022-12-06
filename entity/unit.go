package entity

import "skharv/2DRTS/component"

type Unit struct {
	component.Color
	component.Position
	component.Radius
	component.Selected
}
