package systems

import (
	"game2d/components"
	"game2d/config"
	"game2d/entities"
	"math/rand"
	"reflect"

	"github.com/go-gl/mathgl/mgl32"
)

var possibleObstacles = [][]interface{}{
    // white boar
	{
		&components.PositionComponent{X: config.C.ScreenWidth, Y: 0},
		&components.ObjectComponent{Width: 2 * 38, Height: 2 * 28},
		&components.VelocityComponent{X: -450, Y: 0},
		&components.ColliderComponent{},
		&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{4, 2}, TexCoordsEnd: mgl32.Vec2{43, 32}, TextureID: 5},
		&components.SpriteAnimationComponent{
			Begin:      mgl32.Vec2{4, 2},
			End:        mgl32.Vec2{43, 32},
			Offset:     mgl32.Vec2{48, 0},
			TimeOffset: 1.0 / 10.0,
			CurrTime:   0,
			Length:     6,
		},
	},
	// brown boar
	{
		&components.PositionComponent{X: config.C.ScreenWidth, Y: 0},
		&components.ObjectComponent{Width: 2 * 38, Height: 2 * 28},
		&components.VelocityComponent{X: -500, Y: 0},
		&components.ColliderComponent{},
		&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{4, 66}, TexCoordsEnd: mgl32.Vec2{43, 96}, TextureID: 5},
		&components.SpriteAnimationComponent{
			Begin:      mgl32.Vec2{4, 66},
			End:        mgl32.Vec2{43, 96},
			Offset:     mgl32.Vec2{48, 0},
			TimeOffset: 1.0 / 10.0,
			CurrTime:   0,
			Length:     6,
		},
	},
	// black boar
	{
		&components.PositionComponent{X: config.C.ScreenWidth, Y: 0},
		&components.ObjectComponent{Width: 2 * 38, Height: 2 * 28},
		&components.VelocityComponent{X: -550, Y: 0},
		&components.ColliderComponent{},
		&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{4, 34}, TexCoordsEnd: mgl32.Vec2{43, 64}, TextureID: 5},
		&components.SpriteAnimationComponent{
			Begin:      mgl32.Vec2{4, 34},
			End:        mgl32.Vec2{43, 64},
			Offset:     mgl32.Vec2{48, 0},
			TimeOffset: 1.0 / 10.0,
			CurrTime:   0,
			Length:     6,
		},
	},
	// bee
	{
		&components.PositionComponent{X: config.C.ScreenWidth, Y: 1.2 * config.C.PlayerHeight},
		&components.ObjectComponent{Width: 2 * 24, Height: 2 * 40},
		&components.VelocityComponent{X: -600, Y: 0},
		&components.ColliderComponent{},
		&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{84, 5}, TexCoordsEnd: mgl32.Vec2{120, 44}, TextureID: 6},
		&components.SpriteAnimationComponent{
			Begin:      mgl32.Vec2{22, 5},
			End:        mgl32.Vec2{58, 44},
			Offset:     mgl32.Vec2{64, 0},
			TimeOffset: 1.0 / 20.0,
			CurrTime:   0,
			Length:     4,
		},
	},
	// snail
	{
		&components.PositionComponent{X: config.C.ScreenWidth, Y: 0},
		&components.ObjectComponent{Width: 2 * 36, Height: 2 * 22},
		&components.VelocityComponent{X: -330, Y: 0},
		&components.ColliderComponent{},
		&components.SpriteComponent{TexCoordsBegin: mgl32.Vec2{194, 10}, TexCoordsEnd: mgl32.Vec2{230, 32}, TextureID: 7},
		&components.SpriteAnimationComponent{
			Begin:      mgl32.Vec2{0, 10},
			End:        mgl32.Vec2{39, 32},
			Offset:     mgl32.Vec2{48, 0},
			TimeOffset: 1.0 / 10.0,
			CurrTime:   0,
			Length:     8,
		},
	},
}

type RandomObstacleSystem struct {
	System
	lastObstacleTime float64
}

func (r *RandomObstacleSystem) Update(game *Game, dt float64) {
	if r.shouldGenerateObstacle(dt) {
		obstacleEntity := r.createNewObstacle()
		game.AddEntity(&obstacleEntity.Entity)

	}
}

func Clone(inter interface{}) interface{} {
	nInter := reflect.New(reflect.TypeOf(inter).Elem())

	val := reflect.ValueOf(inter).Elem()
	nVal := nInter.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}

	return nInter.Interface()
}

func (r *RandomObstacleSystem) createNewObstacle() *entities.Obstacle {
	roc := possibleObstacles[rand.Intn(len(possibleObstacles))]
	//deep copy components
	randomObstacleComponents := make([]interface{}, 0, len(roc))
	for _, c := range roc {
		v := Clone(c)
		randomObstacleComponents = append(randomObstacleComponents, v)
	}
	obstacleEntity := &entities.Obstacle{
		Entity: entities.Entity{
			ID:         entities.GenerateUniqueEntityID(),
			Components: randomObstacleComponents,
		},
	}
	return obstacleEntity
}

func (r *RandomObstacleSystem) shouldGenerateObstacle(dt float64) bool {
	r.lastObstacleTime += dt
	if r.lastObstacleTime >= 1*rand.Float64()+1 {
		r.lastObstacleTime = 0
		return true
	}
	return false
}
