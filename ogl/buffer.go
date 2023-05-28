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
func (self *Buffer) Delete() {
	gl.DeleteBuffers(1, &self.bufferID)
	self.bufferID = 0
	self.bufferSize = 0
}
func (self *Buffer) Bind(bindingTarget uint32) {
	if self.currentBinding != 0 {
		self.UnBind()
	}
	self.currentBinding = bindingTarget
	gl.BindBuffer(self.currentBinding, self.bufferID)
}
func (self *Buffer) UnBind() {
	self.currentBinding = 0
	gl.BindBuffer(self.currentBinding, self.bufferID)
}
func (self *Buffer) MapBuffer(bufferTarget uint32, access uint32) unsafe.Pointer {
	return gl.MapBuffer(bufferTarget, access)
}
func (self *Buffer) UnMap(bufferTarget uint32) {
	gl.UnmapBuffer(bufferTarget)
}
func (self *Buffer) BindBufferBase(bufferTarget uint32, bindingIndex uint32) {
	gl.BindBufferBase(bufferTarget, bindingIndex, self.bufferID)
}

func (self *Buffer) createBuffer() {
	gl.GenBuffers(1, &self.bufferID)
}
func allocateBuffer[T any](buffer *Buffer, usage uint32, data *[]T) {
	buffer.Bind(gl.ARRAY_BUFFER)
	var temp T
	buffer.bufferSize = len(*data) * int(reflect.TypeOf(temp).Size())
	gl.BufferData(gl.ARRAY_BUFFER, buffer.bufferSize, gl.Ptr(*data), usage)
	buffer.UnBind()
}
