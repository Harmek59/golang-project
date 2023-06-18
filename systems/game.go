package systems

import (
	"fmt"
	"game2d/components"
	"game2d/config"
	"game2d/entities"

	"github.com/go-gl/mathgl/mgl32"
)

type SystemI interface {
	Update(game *Game, dt float64)
}

type Game struct {
	Entities []*entities.Entity
	IsOver   bool
	Score    int
	systems  []SystemI
}

func NewGame() *Game {
	entities.CreateNewPlayer()
	game := &Game{
		Score:  0,
		IsOver: false,
	}
	// background
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 2, Y: -config.C.ScreenHeight/2 + config.C.CameraYOffset},
				&components.ObjectComponent{
					Width:  config.C.ScreenWidth,
					Height: config.C.ScreenHeight,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{448, 106}, TexCoordsEnd: mgl32.Vec2{768, 286}},
			},
		},
	)
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 2, Y: -config.C.ScreenHeight/2 + config.C.CameraYOffset},
				&components.ObjectComponent{
					Width:  config.C.ScreenWidth,
					Height: config.C.ScreenHeight,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{768, 106}, TexCoordsEnd: mgl32.Vec2{1088, 286}},
			},
		},
	)
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 2, Y: -config.C.ScreenHeight/2 + config.C.CameraYOffset},
				&components.ObjectComponent{
					Width:  config.C.ScreenWidth,
					Height: config.C.ScreenHeight,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{1088, 106}, TexCoordsEnd: mgl32.Vec2{1408, 286}},
			},
		},
	)

	game.AddEntity(&entities.CreateNewPlayer().Entity)
	game.systems = append(game.systems, &InputSystem{})
	game.systems = append(game.systems, &MovementSystem{})
	game.systems = append(game.systems, &CollisionSystem{})
	game.systems = append(game.systems, &SpriteAnimationSystem{})
	game.systems = append(game.systems, &RandomObstacleSystem{})
	game.systems = append(game.systems, CreateRenderSystem())
	return game
}

func (game *Game) AddEntity(entity *entities.Entity) {
	game.Entities = append(game.Entities, entity)
}

func (game *Game) DeleteEntity(entity *entities.Entity) {
	for idx, e := range game.Entities {
		if e.ID == entity.ID {
			game.Entities = append(game.Entities[0:idx], game.Entities[idx+1:]...)
			break
		}
	}

}

func (game *Game) GameOver() {
	fmt.Printf("Game Over\n")

	game.IsOver = true
	game.Score = 0
}

func (game *Game) FindPlayerEntity() *entities.Entity {
	for _, entity := range game.Entities {
		if entity.GetComponent(&components.JumpableComponent{}) != nil {
			return entity
		}
	}
	panic("Player Entity must be defined")
}

func (game *Game) Update(dt float64) error {
	if(game.IsOver){
	    dt = 0.0
	}
	for _, system := range game.systems {
		system.Update(game, dt)
	}
	return nil
}
