package systems

import (
	"fmt"
	"game2d/components"
	"game2d/config"
	"time"
)

type MovementSystem struct{}

func (p *MovementSystem) Update(game *Game) {
	// Update the position of entities based on their velocity
	for _, entity := range game.Entities {
		time.Sleep(1 * time.Second)
		if positionComponent := entity.GetComponent(&components.PositionComponent{}); positionComponent != nil {
			position := positionComponent.(*components.PositionComponent)

			if velocityComponent := entity.GetComponent(&components.VelocityComponent{}); velocityComponent != nil {
				velocity := velocityComponent.(*components.VelocityComponent)

				jumpableComponent := entity.GetComponent(&components.JumpableComponent{})

				if jumpableComponent != nil && jumpableComponent.(*components.JumpableComponent).IsJumping {
					// Apply gravity
					velocity.Y -= config.C.Gravity
				}

				// Handle collision with the ground
				if velocity.Y != 0 {
					if objectComponent := entity.GetComponent(&components.ObjectComponent{}); objectComponent != nil {
						object := objectComponent.(*components.ObjectComponent)
						fmt.Println(velocity.Y, position.Y, object.Height, object.Height+velocity.Y)
						if velocity.Y < 0 && position.Y+velocity.Y < 0 {
							position.Y = 0
							velocity.Y = 0
							fmt.Printf("Entity %d stopped jumping\n", entity.ID)

							if jumpableComponent != nil {
								jumpableComponent.(*components.JumpableComponent).IsJumping = false
							}
						}
					}

				}

				// Update position based on velocity
				position.X += velocity.X
				position.Y += velocity.Y

				//TODO check if < or <=
				if position.X < 0 {
					game.DeleteEntity(entity)
					game.Score += 1
				}

				fmt.Printf("Entity %d position %f %f\n", entity.ID, position.X, position.Y)

			}
		}
	}
}
