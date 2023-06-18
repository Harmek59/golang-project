package systems

import (
	"game2d/components"
	"game2d/core"
	"game2d/ogl"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	triangles = []float32{
		-0.5, 0.5, 0, 0.0, 1.0,
		-0.5, -0.5, 0, 0.0, 0.0,
		0.5, -0.5, 0, 1.0, 0.0,

		0.5, 0.5, 0, 1.0, 1.0,
		0.5, -0.5, 0, 1.0, 0.0,
		-0.5, 0.5, 0, 0.0, 1.0,
	}
)

type RenderSystem struct {
	mainShader ogl.Shader
	vertices   ogl.Vertices
	texture    ogl.Texture
	camera     core.Camera
}

func CreateRenderSystem() *RenderSystem {
	var rs RenderSystem
	var err error
	rs.mainShader, err = ogl.CreateShader("shaders/mainShader.vert", "shaders/mainShader.frag")
	if err != nil {
		panic(err)
	}
	rs.texture, err = ogl.CreateTexture("Texture.png")
	if err != nil {
		panic(err)
	}
	rs.vertices = ogl.CreateVertices(triangles, []uint32{3, 2})
	rs.camera = core.CreateCamera()
	return &rs
}
func (r *RenderSystem) Delete() {
	r.mainShader.Delete()
	r.vertices.Delete()
	r.texture.Delete()
}

func (r *RenderSystem) Update(Game *Game, dt float64) {
	// Render entities on the screen
	r.mainShader.Use()
	gl.ActiveTexture(gl.TEXTURE0)
	r.texture.Bind()
	r.mainShader.SetInt("text", 0)
	r.mainShader.SetMat4("projection", r.camera.GetProjectionMatrix())
	r.mainShader.SetMat4("view", r.camera.GetViewMatrix())
	for _, entity := range Game.Entities {
		positionComponent := entity.GetComponent(&components.PositionComponent{})
		objectComponent := entity.GetComponent(&components.ObjectComponent{})
		spriteComponent := entity.GetComponent(&components.SpriteComponent{})
		if positionComponent != nil && objectComponent != nil {
			position := positionComponent.(*components.PositionComponent)
			size := objectComponent.(*components.ObjectComponent)
			r.mainShader.SetVec2("position", mgl32.Vec2{float32(position.X), float32(position.Y)})
			r.mainShader.SetVec2("scale", mgl32.Vec2{float32(size.Width), float32(size.Height)})
			r.mainShader.SetVec3("color", mgl32.Vec3{1.0, 0.5, 0.4})
			if spriteComponent != nil {
				sprite := spriteComponent.(*components.SpriteComponent)
				r.mainShader.SetVec2("textureBegin", mgl32.Vec2{sprite.TexCoordsBegin[0] / float32(r.texture.Width), sprite.TexCoordsEnd[1] / float32(r.texture.Height)})
				r.mainShader.SetVec2("textureEnd", mgl32.Vec2{sprite.TexCoordsEnd[0] / float32(r.texture.Width), sprite.TexCoordsBegin[1] / float32(r.texture.Height)})
			} else {
				r.mainShader.SetVec2("textureBegin", mgl32.Vec2{0.0, 0.0})
				r.mainShader.SetVec2("textureEnd", mgl32.Vec2{0.0, 0.0})
			}
			r.vertices.Bind()
			gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangles)/5))
		}
	}
}
