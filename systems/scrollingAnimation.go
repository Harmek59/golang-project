package systems

import (
	"game2d/components"
)

type ScrollingAnimationSystem struct {
	System
}

func (o *ScrollingAnimationSystem) Update(game *Game, dt float64) {
	for _, entity := range game.Entities {
		scrollingAnimationComponent := entity.GetComponent(&components.ScrollingAnimationComponent{})
		spriteComponent := entity.GetComponent(&components.SpriteComponent{})
		if scrollingAnimationComponent != nil && spriteComponent != nil {
			sprite := spriteComponent.(*components.SpriteComponent)
			scroll := scrollingAnimationComponent.(*components.ScrollingAnimationComponent)
			sprite.TexCoordsBegin = sprite.TexCoordsBegin.Add(scroll.Speed.Mul(float32(dt)))
			sprite.TexCoordsEnd = sprite.TexCoordsEnd.Add(scroll.Speed.Mul(float32(dt)))
		}
	}
}
