package ogl

import (
	"reflect"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Buffer struct {
	bufferID       uint32
	currentBinding uint32
	bufferSize     int
}

func CreateBuffer[T any](usage uint32, data []T) Buffer {
	var buffer Buffer
	buffer.createBuffer()
	allocateBuffer(&buffer, usage, &data)
	return buffer
}
func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.bufferID)
	buffer.bufferID = 0
	buffer.bufferSize = 0
}
func (buffer *Buffer) Bind(bindingTarget uint32) {
	if buffer.currentBinding != 0 {
		buffer.UnBind()
	}
	buffer.currentBinding = bindingTarget
	gl.BindBuffer(buffer.currentBinding, buffer.bufferID)
}
func (buffer *Buffer) UnBind() {
	buffer.currentBinding = 0
	gl.BindBuffer(buffer.currentBinding, buffer.bufferID)
}
func (buffer *Buffer) MapBuffer(bufferTarget uint32, access uint32) unsafe.Pointer {
	return gl.MapBuffer(bufferTarget, access)
}
func (buffer *Buffer) UnMap(bufferTarget uint32) {
	gl.UnmapBuffer(bufferTarget)
}
func (buffer *Buffer) BindBufferBase(bufferTarget uint32, bindingIndex uint32) {
	gl.BindBufferBase(bufferTarget, bindingIndex, buffer.bufferID)
}

func (buffer *Buffer) createBuffer() {
	gl.GenBuffers(1, &buffer.bufferID)
}
func allocateBuffer[T any](buffer *Buffer, usage uint32, data *[]T) {
	buffer.Bind(gl.ARRAY_BUFFER)
	var temp T
	buffer.bufferSize = len(*data) * int(reflect.TypeOf(temp).Size())
	gl.BufferData(gl.ARRAY_BUFFER, buffer.bufferSize, gl.Ptr(*data), usage)
	buffer.UnBind()
}
