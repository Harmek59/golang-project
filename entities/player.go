package entities

import (
	"game2d/components"
	"game2d/config"

	"github.com/go-gl/mathgl/mgl32"
)

type Player struct {
	Entity
}

func CreateNewPlayer() *Player {
	playerEntity := &Player{
		Entity: Entity{
			ID: GenerateUniqueEntityID(),
			Components: []interface{}{
				&components.PositionComponent{X: -config.C.ScreenWidth / 4, Y: 0},
				&components.JumpableComponent{IsJumping: false},
				&components.ObjectComponent{
					Width:  config.C.PlayerWidth,
					Height: config.C.PlayerHeight,
				},
				&components.VelocityComponent{X: 0, Y: 0},
				&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{16, 136}, TexCoordsEnd: mgl32.Vec2{41, 168}},
				&components.SpriteAnimationComponent{
                    Begin: mgl32.Vec2{16, 134}, 
                    End: mgl32.Vec2{41, 168}, 
                    Offset: mgl32.Vec2{56, 0}, 
                    TimeOffset: 1.0 / 12.0, 
                    CurrTime: 0, 
                    Length: 8,
                },
			},
		},
	}
	return playerEntity
}
