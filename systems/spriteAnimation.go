package systems

import (
	"game2d/components"
	"math"
)

type SpriteAnimationSystem struct {
	System
}

func (o *SpriteAnimationSystem) Update(game *Game, dt float64) {
	for _, entity := range game.Entities {
		spriteAnimationComponent := entity.GetComponent(&components.SpriteAnimationComponent{})
		spriteComponent := entity.GetComponent(&components.SpriteComponent{})
		if spriteAnimationComponent != nil && spriteComponent != nil {
			sprite := spriteComponent.(*components.SpriteComponent)
			spriteAnim := spriteAnimationComponent.(*components.SpriteAnimationComponent)
			spriteAnim.CurrTime = float32(math.Mod(float64(spriteAnim.CurrTime)+dt, float64(spriteAnim.TimeOffset*float32(spriteAnim.Length))))
			currentFrame := int(spriteAnim.CurrTime / spriteAnim.TimeOffset)
			sprite.TexCoordsBegin = spriteAnim.Begin.Add(spriteAnim.Offset.Mul(float32(currentFrame)))
			sprite.TexCoordsEnd = spriteAnim.End.Add(spriteAnim.Offset.Mul(float32(currentFrame)))
		}
	}
}
