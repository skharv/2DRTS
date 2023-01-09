package entity

import "skharv/2DRTS/component"

type WalkableSpace struct {
	component.Color
	component.NavMesh
	component.Triangle
}
