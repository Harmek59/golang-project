package components

import (
	"github.com/go-gl/mathgl/mgl32"
)

type SpriteComponent struct {
	TexCoordsBegin mgl32.Vec2
	TexCoordsEnd   mgl32.Vec2
	TextureID      int
}
