package systems

import "fmt"

type CollisionSystem struct {
	System
}

func (o *CollisionSystem) Update(game *Game) {
	playerEntity := game.FindPlayerEntity()

	for _, entity := range game.Entities {
		if playerEntity.ID == entity.ID {
			continue
		}

		if playerEntity.IsColliding(entity) {
			fmt.Println("Collision detected!")
			// Perform additional game over logic here
			game.GameOver()
			return
		}

	}
}
