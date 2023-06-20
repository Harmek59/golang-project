package systems

import (
	"game2d/components"
)

type CollisionSystem struct {
	System
}

func (o *CollisionSystem) Update(game *Game, dt float64) {
	if dt == 0.0 {
		return
	}
	playerEntity := game.FindPlayerEntity()

	for _, entity := range game.Entities {
		colliderComponent := entity.GetComponent(&components.ColliderComponent{})
		if colliderComponent != nil {
			if playerEntity.IsColliding(entity) {
				// Perform additional game over logic here
				game.GameOver()
				return
			}
		}
	}
}
