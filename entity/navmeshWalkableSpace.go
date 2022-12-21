package entity

import "skharv/2DRTS/component"

type NavmeshWalkableSpace struct {
	component.Color
	component.NavMesh
	component.Rectangle
}
