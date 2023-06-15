package entities

import (
	"game2d/components"
	"game2d/config"
)

type Player struct {
	Entity
}

func CreateNewPlayer() *Player {
	playerEntity := &Player{
		Entity: Entity{
			ID: GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: 0, Y: 0},
				&components.JumpableComponent{IsJumping: false},
				&components.ObjectComponent{
					Width:  config.C.PlayerWidth,
					Height: config.C.PlayerHeight,
				},
				&components.VelocityComponent{X: 0, Y: 0},
			},
		},
	}
	return playerEntity
}
