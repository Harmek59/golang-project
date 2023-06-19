package ogl

import (
	"reflect"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type InstancedVertices struct {
	verticesBuffer          Buffer
	instanceBuffer          Buffer
	vao                     uint32
	numberOfModelAttributes int
}

func CreateInstancedVertices(vertArr []float32, attribSizes []uint32) InstancedVertices {
	var vertices InstancedVertices
	gl.GenVertexArrays(1, &vertices.vao)

	vertices.Bind()

	vertices.verticesBuffer = CreateBuffer(gl.STATIC_DRAW, vertArr)
	vertices.verticesBuffer.Bind(gl.ARRAY_BUFFER)

	vertices.UnBind()

	vertices.setAttributeSizes(attribSizes)

	return vertices
}
func (vertices *InstancedVertices) Delete() {
	gl.DeleteVertexArrays(1, &vertices.vao)
	vertices.verticesBuffer.Delete()
	vertices.instanceBuffer.Delete()
	vertices.vao = 0
}
func (vertices *InstancedVertices) Bind() {
	gl.BindVertexArray(vertices.vao)
}
func (vertices *InstancedVertices) UnBind() {
	gl.BindVertexArray(0)
}
func (vertices *InstancedVertices) GetInstanceBuffer() Buffer {
	return vertices.instanceBuffer
}
func (vertices *InstancedVertices) SetUpInstanceBuffer(bufferSizeByte int, attribSizes []uint32, divisorSizes []uint32) {
	if len(divisorSizes) != len(attribSizes) {
		panic("InstancedVertices: attribSizes and divisorSizes must be equal")
	}

	var offsetOfAllAttributes uint32
	for _, v := range attribSizes {
		offsetOfAllAttributes += v
	}

	vertices.instanceBuffer = CreateBuffer(gl.DYNAMIC_DRAW, make([]byte, bufferSizeByte))
	vertices.Bind()
	vertices.instanceBuffer.Bind(gl.ARRAY_BUFFER)

	var beginOfAttribute uint32
	var positionOfAttribute = uint32(vertices.numberOfModelAttributes)
	sizeOfFloat := reflect.TypeOf(float32(0.0)).Size()
	for _, v := range attribSizes {
		gl.EnableVertexAttribArray(positionOfAttribute)
		gl.VertexAttribPointer(positionOfAttribute, int32(v), gl.FLOAT, false,
			int32(offsetOfAllAttributes*uint32(sizeOfFloat)), gl.PtrOffset(int(beginOfAttribute)*int(sizeOfFloat)))
		gl.VertexAttribDivisor(positionOfAttribute, divisorSizes[positionOfAttribute-uint32(vertices.numberOfModelAttributes)])

		beginOfAttribute += v
		positionOfAttribute++
	}

	vertices.verticesBuffer.UnBind()
	vertices.UnBind()
}
func (vertices *InstancedVertices) setAttributeSizes(attribSizes []uint32) {
	vertices.Bind()
	vertices.numberOfModelAttributes = len(attribSizes)
	vertices.verticesBuffer.Bind(gl.ARRAY_BUFFER)

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

	vertices.verticesBuffer.UnBind()
	vertices.UnBind()
}
