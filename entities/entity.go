package entities

import (
	"fmt"
	"game2d/components"
	"image"
)

var (
	entityCounter int
)

type Entity struct {
	ID         int
	Components []interface{}
}

func GenerateUniqueEntityID() int {
	entityCounter++
	return entityCounter
}

func (entity *Entity) GetComponent(componentType interface{}) interface{} {
	for _, component := range entity.Components {
		if fmt.Sprintf("%T", component) == fmt.Sprintf("%T", componentType) {
			return component
		}
	}
	return nil
}

func (entity *Entity) constructRect() *image.Rectangle {
	var entityPositionComponent *components.PositionComponent
	var entityobjectComponent *components.ObjectComponent
	if positionComponent := entity.GetComponent(&components.PositionComponent{}); positionComponent != nil {
		entityPositionComponent = positionComponent.(*components.PositionComponent)
	} else {
		return nil
	}
	if objectComponent := entity.GetComponent(&components.ObjectComponent{}); objectComponent != nil {
		entityobjectComponent = objectComponent.(*components.ObjectComponent)
	} else {
		return nil
	}

	entityRect := image.Rect(
		int(entityPositionComponent.X+entityobjectComponent.Width/8),
		int(entityPositionComponent.Y),
		int(entityPositionComponent.X+0.7*entityobjectComponent.Width),
		int(entityPositionComponent.Y+0.8*entityobjectComponent.Height),
	)
	return &entityRect
}

func (entity *Entity) IsColliding(obstacle *Entity) bool {

	if entityRect := entity.constructRect(); entityRect != nil {
		if obstacleRect := obstacle.constructRect(); obstacleRect != nil {
			collision := entityRect.Overlaps(*obstacleRect)
			return collision
		}
	}
	return false
}
