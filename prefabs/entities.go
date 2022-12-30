package prefabs

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/entity"
	"skharv/2DRTS/helper/globals"
)

type Entities struct {
	Units   map[string]entity.Unit
	counter int
}

func Initialise() *Entities {
	e := Entities{map[string]entity.Unit{}, 0}

	unitA := entity.Unit{
		Color:    component.NewColor(255, 255, 255, 255),
		Facing:   component.NewFacing(0),
		Id:       component.NewId(0),
		Owner:    component.NewOwner(0),
		Position: component.NewPosition(0, 0),
		Radius:   component.NewRadius(10),
		Selected: component.NewSelected(false),
		Speed:    component.NewSpeed(2),
		State:    component.NewState(globals.Idle),
		Target:   component.NewTarget(0, 0),
		TurnRate: component.NewTurnRate(10),
		Weight:   component.NewWeight(50),
	}
	e.Units["unitA"] = unitA

	unitB := entity.Unit{
		Color:    component.NewColor(255, 255, 255, 255),
		Facing:   component.NewFacing(0),
		Id:       component.NewId(0),
		Owner:    component.NewOwner(0),
		Position: component.NewPosition(0, 0),
		Radius:   component.NewRadius(20),
		Selected: component.NewSelected(false),
		Speed:    component.NewSpeed(1),
		State:    component.NewState(globals.Idle),
		Target:   component.NewTarget(0, 0),
		TurnRate: component.NewTurnRate(5),
		Weight:   component.NewWeight(150),
	}
	e.Units["unitB"] = unitB

	return &e
}

func (e *Entities) SpawnUnit(name string, posX, posY float64, owner int) *entity.Unit {
	unit := e.Units[name]

	unit.Id = component.NewId(e.counter)
	e.counter++

	unit.Position = component.NewPosition(posX, posY)

	unit.Owner = component.NewOwner(owner)

	return &unit
}
