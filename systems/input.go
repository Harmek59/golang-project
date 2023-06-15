package systems

import (
	"fmt"
	"game2d/components"
	"game2d/config"
)

type InputSystem struct {
	System
}

func (i *InputSystem) Update(game *Game) {
	playerEntity := game.FindPlayerEntity()
	playerJumpableComponent := playerEntity.GetComponent(&components.JumpableComponent{}).(*components.JumpableComponent)

	// TODO add detecting space input

	if !playerJumpableComponent.IsJumping {
		fmt.Printf("Entity %d started jumping\n", playerEntity.ID)
		playerJumpableComponent.IsJumping = true

		velocityComponent := playerEntity.GetComponent(&components.VelocityComponent{}).(*components.VelocityComponent)
		velocityComponent.Y += config.C.JumpForce

	}
}
