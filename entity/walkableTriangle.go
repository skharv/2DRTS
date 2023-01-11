package entity

import "skharv/2DRTS/component"

type WalkableTriangle struct {
	component.Color
	component.NavMesh
	component.Triangle
}
