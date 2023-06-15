package ogl

import (
	"reflect"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Vertices struct {
	buffer Buffer
	vao    uint32
}

func CreateVertices(vertArr []float32, attribSizes []uint32) Vertices {
	var vertices Vertices
	gl.GenVertexArrays(1, &vertices.vao)

	vertices.Bind()

	vertices.buffer = CreateBuffer(gl.STATIC_DRAW, vertArr)
	vertices.buffer.Bind(gl.ARRAY_BUFFER)

	vertices.UnBind()

	vertices.setAttributeSizes(attribSizes)

	return vertices
}
func (vertices *Vertices) Delete() {
	gl.DeleteVertexArrays(1, &vertices.vao)
	vertices.buffer.Delete()
	vertices.vao = 0
}
func (vertices *Vertices) Bind() {
	gl.BindVertexArray(vertices.vao)
}
func (vertices *Vertices) UnBind() {
	gl.BindVertexArray(0)
}

func (vertices *Vertices) setAttributeSizes(attribSizes []uint32) {
	vertices.Bind()
	vertices.buffer.Bind(gl.ARRAY_BUFFER)

	var offsetOfAllAttributes uint32
	for _, v := range attribSizes {
		offsetOfAllAttributes += v
	}

	var beginOfAttribute uint32
	var positionOfAttribute uint32
	sizeOfFloat := reflect.TypeOf(float32(0.0)).Size()
	for _, v := range attribSizes {
		gl.EnableVertexAttribArray(positionOfAttribute)
		gl.VertexAttribPointer(positionOfAttribute, int32(v), gl.FLOAT, false,
			int32(offsetOfAllAttributes*uint32(sizeOfFloat)), gl.PtrOffset(int(beginOfAttribute)*int(sizeOfFloat)))

		beginOfAttribute += v
		positionOfAttribute++
	}

	vertices.buffer.UnBind()
	vertices.UnBind()
}
