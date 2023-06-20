package systems

import (
	"fmt"
	"game2d/components"
	"game2d/config"
)

type MovementSystem struct{}

func (p *MovementSystem) Update(game *Game, dt float64) {
	// Update the position of entities based on their velocity
	for _, entity := range game.Entities {
		if positionComponent := entity.GetComponent(&components.PositionComponent{}); positionComponent != nil {
			position := positionComponent.(*components.PositionComponent)

			if velocityComponent := entity.GetComponent(&components.VelocityComponent{}); velocityComponent != nil {
				velocity := velocityComponent.(*components.VelocityComponent)

				jumpableComponent := entity.GetComponent(&components.JumpableComponent{})

				if jumpableComponent != nil && jumpableComponent.(*components.JumpableComponent).IsJumping {
					// Apply gravity
					velocity.Y -= dt * config.C.Gravity
				}

				// Update position based on velocity
				position.X += dt * velocity.X
				position.Y += dt * velocity.Y

				// Handle collision with the ground
				if velocity.Y != 0 {
					if objectComponent := entity.GetComponent(&components.ObjectComponent{}); objectComponent != nil {
						if position.Y < 0 {
							position.Y = 0
							velocity.Y = 0

							if jumpableComponent != nil {
								jumpableComponent.(*components.JumpableComponent).IsJumping = false
							}
						}
					}

				}

				if position.X < -1000 {
					game.DeleteEntity(entity)
					fmt.Println("DESPAWN")
					game.Score += 1
				}
			}
		}
	}
}
