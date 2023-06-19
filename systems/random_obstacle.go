package systems

import (
	"game2d/components"
	"game2d/config"
	"game2d/entities"
	"math/rand"
)

type RandomObstacleSystem struct {
	System
	lastObstacleTime float64
}

func (r *RandomObstacleSystem) Update(game *Game, dt float64) {
	if r.shouldGenerateObstacle(dt) {
		obstacleEntity := r.createNewObstacle()
		game.AddEntity(&obstacleEntity.Entity)

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
				&components.ColliderComponent{},
			},
		},
	}
	return obstacleEntity
}

func (r *RandomObstacleSystem) shouldGenerateObstacle(dt float64) bool {
	r.lastObstacleTime += dt
	if r.lastObstacleTime >= 1*rand.Float64()+1 {
		r.lastObstacleTime = 0
		return true
	}
	return false
}

func (r *RandomObstacleSystem) generateRandomYPosition() float64 {
	if rand.Float64() > 0.6 {
		return 0
	} else {
		return 1.2 * config.C.PlayerHeight
	}
}

func (r *RandomObstacleSystem) generateRandomWidth() float64 {
	return rand.Float64()*30 + 30
}

func (r *RandomObstacleSystem) generateRandomHeight() float64 {
	return rand.Float64()*30 + 30
}
