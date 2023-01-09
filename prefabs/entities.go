package prefabs

import (
	"skharv/2DRTS/component"
	"skharv/2DRTS/entity"
	"skharv/2DRTS/helper/globals"
)

type Entities struct {
	counter int
}

func (e *Entities) SpawnUnit(unit string, posX, posY, player int) *entity.Unit {
	var radius = component.Radius{}
	var speed = component.Speed{}
	var turnRate = component.TurnRate{}
	var weight = component.Weight{}

	switch unit {
	case "unitA":
		radius = component.NewRadius(10)
		speed = component.NewSpeed(2)
		turnRate = component.NewTurnRate(10)
		weight = component.NewWeight(50)
	case "unitB":
		radius = component.NewRadius(20)
		speed = component.NewSpeed(1)
		turnRate = component.NewTurnRate(5)
		weight = component.NewWeight(150)
	case "unitC":
		radius = component.NewRadius(5)
		speed = component.NewSpeed(1.5)
		turnRate = component.NewTurnRate(15)
		weight = component.NewWeight(300)
	case "unitD":
		radius = component.NewRadius(12)
		speed = component.NewSpeed(0.75)
		turnRate = component.NewTurnRate(20)
		weight = component.NewWeight(75)
	}

	chunk := component.NewChunk(0, 0)
	color := component.NewColor(255, 255, 255, 255)
	facing := component.NewFacing(0)
	id := component.NewId(e.counter)
	name := component.NewName(unit)
	owner := component.NewOwner(player)
	position := component.NewPosition(float64(posX), float64(posY))
	selected := component.NewSelected(false)
	state := component.NewState(globals.Idle)
	target := component.NewTarget(0, 0)
	velocity := component.NewVelocity(0, 0)

	entity := entity.Unit{
		Chunk:    chunk,
		Color:    color,
		Facing:   facing,
		Id:       id,
		Name:     name,
		Owner:    owner,
		Position: position,
		Radius:   radius,
		Selected: selected,
		Speed:    speed,
		State:    state,
		Target:   target,
		TurnRate: turnRate,
		Velocity: velocity,
		Weight:   weight,
	}

	e.counter++

	return &entity
}

func (e *Entities) SpawnBuilding(name string, posX, posY, player int) *entity.Building {
	var radius = component.Radius{}

	switch name {
	case "buildingA":
		radius = component.NewRadius(40)
	case "buildingB":
		radius = component.NewRadius(75)
	case "buildingC":
		radius = component.NewRadius(50)
	case "buildingD":
		radius = component.NewRadius(120)
	}

	chunk := component.NewChunk(0, 0)
	color := component.NewColor(255, 255, 255, 255)
	id := component.NewId(e.counter)
	owner := component.NewOwner(player)
	position := component.NewPosition(float64(posX), float64(posY))
	selected := component.NewSelected(false)
	state := component.NewState(globals.Idle)
	target := component.NewTarget(0, 0)

	building := entity.Building{
		Chunk:    chunk,
		Color:    color,
		Id:       id,
		Owner:    owner,
		Position: position,
		Radius:   radius,
		Selected: selected,
		State:    state,
		Target:   target,
	}

	e.counter++

	return &building
}
