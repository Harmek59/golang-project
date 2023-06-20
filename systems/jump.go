package systems

import (
	"game2d/components"
	"game2d/config"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type JumpSystem struct {
	System
}

func (i *JumpSystem) Update(game *Game, _ float64) {
	playerEntity := game.FindPlayerEntity()
	playerJumpableComponent := playerEntity.GetComponent(&components.JumpableComponent{}).(*components.JumpableComponent)

	if glfw.GetCurrentContext().GetKey(glfw.KeySpace) == glfw.Press && !playerJumpableComponent.IsJumping {
		playerJumpableComponent.IsJumping = true

		velocityComponent := playerEntity.GetComponent(&components.VelocityComponent{}).(*components.VelocityComponent)
		velocityComponent.Y += config.C.JumpForce

	}
}
