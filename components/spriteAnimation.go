package components

import (
	"github.com/go-gl/mathgl/mgl32"
)

type SpriteAnimationComponent struct {
	Begin      mgl32.Vec2
	End        mgl32.Vec2
	Offset     mgl32.Vec2
	TimeOffset float32
	CurrTime   float32
	Length     int
}
