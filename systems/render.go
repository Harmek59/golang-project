package systems

import (
	"game2d/components"
	"game2d/core"
	"game2d/ogl"
	"strconv"

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
	mainShader  ogl.Shader
	vertices    ogl.InstancedVertices
	textures    []ogl.Texture
	camera      core.Camera
	maxEntities int
}

func CreateRenderSystem() *RenderSystem {
	var rs RenderSystem
	var err error
	rs.maxEntities = 100
	rs.mainShader, err = ogl.CreateShader("shaders/mainShader.vert", "shaders/mainShader.frag")
	if err != nil {
		panic(err)
	}
	addTexture := func(fileName string) {
		texture, err := ogl.CreateTexture(fileName)
		if err != nil {
			panic(err)
		}
		rs.textures = append(rs.textures, texture)
	}
	addTexture("character.png")
	addTexture("background_layer_1.png")
	addTexture("background_layer_2.png")
	addTexture("background_layer_3.png")
	addTexture("path.png")
	rs.vertices = ogl.CreateInstancedVertices(triangles, []uint32{3, 2})
	rs.vertices.SetUpInstanceBuffer(rs.maxEntities*9*4, []uint32{2, 2, 2, 2, 1}, []uint32{1, 1, 1, 1, 1})
	rs.camera = core.CreateCamera()
	return &rs
}
func (r *RenderSystem) Delete() {
	r.mainShader.Delete()
	r.vertices.Delete()
	for _, t := range r.textures {
		t.Delete()
	}
}
func (r *RenderSystem) Update(Game *Game, _ float64) {
	// Render entities on the screen
	r.mainShader.Use()
	for i, t := range r.textures {
		gl.ActiveTexture(gl.TEXTURE0 + uint32(i))
		t.Bind()
		r.mainShader.SetInt("texture"+strconv.Itoa(i), int32(i))
	}
	r.mainShader.SetMat4("projection", r.camera.GetProjectionMatrix())
	r.mainShader.SetMat4("view", r.camera.GetViewMatrix())
	r.mainShader.SetVec3("color", mgl32.Vec3{1.0, 0.5, 0.4})

	var numOfEntitiesToDraw int32

	buffer := r.vertices.GetInstanceBuffer()
	buffer.Bind(gl.ARRAY_BUFFER)

	ptr := buffer.MapBuffer(gl.ARRAY_BUFFER, gl.WRITE_ONLY|gl.MAP_INVALIDATE_BUFFER_BIT)
	if ptr == nil {
		panic("Failed to map buffer")
	}
	length := r.maxEntities * 9
	slice := (*[1 << 30]float32)(ptr)[:length:length]
	for i, entity := range Game.Entities {
		positionComponent := entity.GetComponent(&components.PositionComponent{})
		objectComponent := entity.GetComponent(&components.ObjectComponent{})
		spriteComponent := entity.GetComponent(&components.SpriteComponent{})
		if positionComponent != nil && objectComponent != nil {
			position := positionComponent.(*components.PositionComponent)
			size := objectComponent.(*components.ObjectComponent)
			slice[i*9+0] = float32(position.X)
			slice[i*9+1] = float32(position.Y)
			slice[i*9+2] = float32(size.Width)
			slice[i*9+3] = float32(size.Height)
			if spriteComponent != nil {
				sprite := spriteComponent.(*components.SpriteComponent)
				begin := mgl32.Vec2{sprite.TexCoordsBegin[0] / float32(r.textures[sprite.TextureID].Width), sprite.TexCoordsEnd[1] / float32(r.textures[sprite.TextureID].Height)}
				end := mgl32.Vec2{sprite.TexCoordsEnd[0] / float32(r.textures[sprite.TextureID].Width), sprite.TexCoordsBegin[1] / float32(r.textures[sprite.TextureID].Height)}
				slice[i*9+4] = float32(begin[0])
				slice[i*9+5] = float32(begin[1])
				slice[i*9+6] = float32(end[0])
				slice[i*9+7] = float32(end[1])
				slice[i*9+8] = float32(sprite.TextureID)
			} else {
				slice[i*9+4] = 0.0
				slice[i*9+5] = 0.0
				slice[i*9+6] = 0.0
				slice[i*9+7] = 0.0
				slice[i*9+8] = -1.0
			}
			numOfEntitiesToDraw++
		}
	}
	buffer.UnMap(gl.ARRAY_BUFFER)
	buffer.UnBind()

	r.vertices.Bind()
	gl.DrawArraysInstanced(gl.TRIANGLES, 0, int32(len(triangles)/5), numOfEntitiesToDraw)
}
