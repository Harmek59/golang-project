package systems

import (
	"game2d/components"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

type DigitOfScoreSystem struct {
	System
}

func (o *DigitOfScoreSystem) Update(game *Game, dt float64) {
	for _, entity := range game.Entities {
		digitOfScoreComponent := entity.GetComponent(&components.DigitOfScoreComponent{})
		spriteComponent := entity.GetComponent(&components.SpriteComponent{})
		if digitOfScoreComponent != nil && spriteComponent != nil {
			sprite := spriteComponent.(*components.SpriteComponent)
			digit := digitOfScoreComponent.(*components.DigitOfScoreComponent)
			v := game.Score / int(math.Pow(10, float64(digit.Digit-1)))
			v = v % 10
			sprite.TexCoordsBegin = mgl32.Vec2{float32((v - 1) * 32), 0}
			sprite.TexCoordsEnd = mgl32.Vec2{float32(v * 32), 32}
		}
	}
}
