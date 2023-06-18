package systems

import (
	"fmt"
	"game2d/components"
	"game2d/config"
	"game2d/entities"
	"math/rand"
	"time"
)

type RandomObstacleSystem struct {
	System
	lastObstacleTime time.Time
}

func (r *RandomObstacleSystem) Update(game *Game, dt float64) {
	if r.shouldGenerateObstacle() {
		// Create a new obstacle entity
		obstacleEntity := r.createNewObstacle()
		game.AddEntity(&obstacleEntity.Entity)
		fmt.Printf("Entity Obstacle %d generated\n", obstacleEntity.ID)

	}
}

func (r *RandomObstacleSystem) createNewObstacle() *entities.Obstacle {
	obstacleEntity := &entities.Obstacle{
		Entity: entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: config.C.ScreenWidth, Y: r.generateRandomYPosition()},
				&components.ObjectComponent{Width: r.generateRandomWidth(), Height: r.generateRandomHeight()},
				&components.VelocityComponent{X: -500, Y: 0},
			},
		},
	}
	return obstacleEntity
}

func (r *RandomObstacleSystem) shouldGenerateObstacle() bool {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(r.lastObstacleTime)
	if elapsedTime >= time.Second*2 {
		r.lastObstacleTime = currentTime
		return true
	}
	return false
}

// TODO implement after obstacle visualization

func (r *RandomObstacleSystem) generateRandomYPosition() float64 {
	return 0
	//return rand.Float64() * (config.C.ScreenHeight - config.C.PlayerHeight)
}

func (r *RandomObstacleSystem) generateRandomWidth() float64 {
	return rand.Float64()*50 + 20
}

func (r *RandomObstacleSystem) generateRandomHeight() float64 {
	return rand.Float64()*50 + 20
}
