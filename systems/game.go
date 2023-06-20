package systems

import (
	"game2d/components"
	"game2d/config"
	"game2d/entities"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"sync"
)

type SystemI interface {
	Update(game *Game, dt float64)
}

type Game struct {
	Entities []*entities.Entity
	Score    int
	mutex    sync.Mutex
	gameover bool
	systems  []SystemI
}

func NewGame() *Game {
	game := &Game{
		Score:    0,
		gameover: false,
	}
	game.mutex.Lock()
	defer game.mutex.Unlock()
	game.setupScene()
	game.setupSystems()
	return game
}

func (game *Game) setupSystems() {
	game.systems = append(game.systems, &JumpSystem{})
	game.systems = append(game.systems, &MovementSystem{})
	game.systems = append(game.systems, &CollisionSystem{})
	game.systems = append(game.systems, &DigitOfScoreSystem{})
	game.systems = append(game.systems, &SpriteAnimationSystem{})
	game.systems = append(game.systems, &ScrollingAnimationSystem{})
	game.systems = append(game.systems, &RandomObstacleSystem{})
	game.systems = append(game.systems, CreateRenderSystem())
}

func (game *Game) setupScene() {
	game.setupBackground()
	game.setupScoreSprites()

	game.AddEntity(&entities.CreateNewPlayer().Entity)
}

func (game *Game) setupScoreSprites() {
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -96 * 1.5, Y: -config.C.ScreenHeight/6 + config.C.CameraYOffset*2},
				&components.ObjectComponent{
					Width:  96,
					Height: 96,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{32, 32}, TextureID: 9},
				&components.DigitOfScoreComponent{Digit: 3},
			},
		},
	)
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -96 * 0.5, Y: -config.C.ScreenHeight/6 + config.C.CameraYOffset*2},
				&components.ObjectComponent{
					Width:  96,
					Height: 96,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{32, 32}, TextureID: 9},
				&components.DigitOfScoreComponent{Digit: 2},
			},
		},
	)
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: 96 * 0.5, Y: -config.C.ScreenHeight/6 + config.C.CameraYOffset*2},
				&components.ObjectComponent{
					Width:  96,
					Height: 96,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{0, 0}, TexCoordsEnd: mgl32.Vec2{32, 32}, TextureID: 9},
				&components.DigitOfScoreComponent{Digit: 1},
			},
		},
	)
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
	game.gameover = true
	game.AddEntity(
		&entities.Entity{
			ID: entities.GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 4, Y: -config.C.ScreenHeight/4 + config.C.CameraYOffset},
				&components.ObjectComponent{
					Width:  422,
					Height: 182,
				},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{1, 0}, TexCoordsEnd: mgl32.Vec2{422, 182}, TextureID: 8},
			},
		},
	)
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
	game.mutex.Lock()
	defer game.mutex.Unlock()
	if game.gameover {
		dt = 0.0
		if glfw.GetCurrentContext().GetKey(glfw.KeyEnter) == glfw.Press {
			playerEntity := game.FindPlayerEntity()
			game.Entities = nil
			game.Score = 0
			game.setupBackground()
			game.setupScoreSprites()
			game.AddEntity(playerEntity)
			game.gameover = false
		}
	}
	for _, system := range game.systems {
		system.Update(game, dt)
	}
	return nil
}
