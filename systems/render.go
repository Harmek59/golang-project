package systems

import (
	"game2d/components"
	"game2d/config"
	"game2d/core"
	"game2d/ogl"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	triangles = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		0.5, 0.5, 0,
		0.5, -0.5, 0,
		-0.5, 0.5, 0,
	}
)

type RenderSystem struct {
	mainShader ogl.Shader
	vertices   ogl.Vertices
	camera     core.Camera
}

func CreateRenderSystem() *RenderSystem {
	var rs RenderSystem
	var err error
	rs.mainShader, err = ogl.CreateShader("shaders/mainShader.vert", "shaders/mainShader.frag")
	if err != nil {
		panic(err)
	}
	rs.vertices = ogl.CreateVertices(triangles, []uint32{3})
	rs.camera = core.CreateCamera(float32(config.C.ScreenWidth), float32(config.C.ScreenHeight))
    return &rs
}
func (r *RenderSystem) Delete() {
	r.mainShader.Delete()
	r.vertices.Delete()
}

var color mgl32.Vec3 = mgl32.Vec3{1.0, 0.5, 0.4}

func (r *RenderSystem) Update(Game *Game, dt float64) {
	// Render entities on the screen
	r.mainShader.Use()
	r.mainShader.SetMat4("projection", r.camera.GetProjectionMatrix())
	r.mainShader.SetMat4("view", r.camera.GetViewMatrix())
	for _, entity := range Game.Entities {
		positionComponent := entity.GetComponent(&components.PositionComponent{})
		objectComponent := entity.GetComponent(&components.ObjectComponent{})
		if positionComponent != nil && objectComponent != nil {
			position := positionComponent.(*components.PositionComponent)
			size := objectComponent.(*components.ObjectComponent)
			r.mainShader.SetVec2("position", mgl32.Vec2{float32(position.X), float32(position.Y)})
			r.mainShader.SetVec2("scale", mgl32.Vec2{float32(size.Width), float32(size.Height)})
			r.mainShader.SetVec3("color", color)
			r.vertices.Bind()
			gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangles)/3))
			continue
			//position := positionComponent.(*components.PositionComponent)
			// Render the entity at its position
			// Replace this with your rendering code
		}
	}
}
