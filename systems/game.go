package systems

import (
	"fmt"
	"game2d/components"
	"game2d/entities"
)

type Game struct {
	Entities []*entities.Entity
	IsOver   bool
	Score    int
}

func NewGame() *Game {
	entities.CreateNewPlayer()
	game := &Game{
		Score:  0,
		IsOver: false,
	}
	game.AddEntity(&entities.CreateNewPlayer().Entity)
	return game
}

func (game *Game) AddEntity(entity *entities.Entity) {
	game.Entities = append(game.Entities, entity)
}

func (game *Game) DeleteEntity(entity *entities.Entity) {
	for idx, e := range game.Entities {
		if e.ID == entity.ID {
			game.Entities = append(game.Entities[0:idx], game.Entities[idx+1:]...)
			fmt.Printf("Deleted entity %d\n", entity.ID)
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

// TODO verify order of systems

func GetSystems() []System {
	return []System{
		&RenderSystem{},
		&InputSystem{},
		&MovementSystem{},
		&CollisionSystem{},
		&RandomObstacleSystem{},
	}
}

func (game *Game) Update() error {
	for _, system := range GetSystems() {
		system.Update(game)
	}
	return nil
}
