package entity

import "skharv/2DRTS/component"

type Cursor struct {
	component.Clicked
	component.Color
	component.Rectangle
}
