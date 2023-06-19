package systems

import (
	"game2d/components"
	"game2d/config"
	"game2d/entities"
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

type SystemI interface {
	Update(game *Game, dt float64)
}

type Game struct {
	Entities []*entities.Entity
	Score    int
	systems  []SystemI
}

func NewGame() *Game {
	game := &Game{
		Score: 0,
	}
	game.setupScene()
	game.setupSystems()
	return game
}

func (game *Game) setupSystems() {
	game.systems = append(game.systems, &InputSystem{})
	game.systems = append(game.systems, &MovementSystem{})
	game.systems = append(game.systems, &CollisionSystem{})
	game.systems = append(game.systems, &SpriteAnimationSystem{})
	game.systems = append(game.systems, &ScrollingAnimationSystem{})
	game.systems = append(game.systems, &RandomObstacleSystem{})
	game.systems = append(game.systems, CreateRenderSystem())
}

func (game *Game) setupScene() {
	game.setupBackground()

	game.AddEntity(&entities.CreateNewPlayer().Entity)
}

func (game *Game) setupBackground() {
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 2, Y: -config.C.ScreenHeight/2 + config.C.CameraYOffset},
				&components.ObjectComponent{
					Width:  config.C.ScreenWidth,
					Height: config.C.ScreenHeight,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{320, 180}, TextureID: 1},
				&components.ScrollingAnimationComponent{Speed: mgl32.Vec2{10, 0}},
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
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{320, 180}, TextureID: 2},
				&components.ScrollingAnimationComponent{Speed: mgl32.Vec2{30, 0}},
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
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{320, 180}, TextureID: 3},
				&components.ScrollingAnimationComponent{Speed: mgl32.Vec2{60, 0}},
			},
		},
	)
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 2, Y: -50.0},
				&components.ObjectComponent{
					Width:  64 * 64,
					Height: 64,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{64 * 32, 31}, TextureID: 4},
				&components.ScrollingAnimationComponent{Speed: mgl32.Vec2{150, 0}},
			},
		},
	)
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
	//TODO add "GAME OVER" display text
	time.Sleep(2 * time.Second)
	playerEntity := game.FindPlayerEntity()
	game.Entities = nil
	game.Score = 0
	game.setupBackground()
	game.AddEntity(playerEntity)

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
	for _, system := range game.systems {
		system.Update(game, dt)
	}
	return nil
}
