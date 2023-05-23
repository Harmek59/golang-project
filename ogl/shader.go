package ogl

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	programID uint32
}

func CreateShader(vertexShaderPath string, fragmentShaderPath string) (Shader, error) {
	var s Shader
	s.createProgram()

	vertShaderCode, err := os.ReadFile(vertexShaderPath)
	if err != nil {
		return s, fmt.Errorf("Failed to load file: %v, err: %v", vertexShaderPath, err)
	}
	shader, err := createAndCompileShader(string(vertShaderCode), gl.VERTEX_SHADER)
	if err != nil {
		return s, fmt.Errorf("%v: %v", vertexShaderPath, err)
	}
	s.attachShaderToProgram(shader)

	fragShaderCode, err := os.ReadFile(fragmentShaderPath)
	if err != nil {
		return s, fmt.Errorf("Failed to load file: %v, err: %v", vertexShaderPath, err)
	}
	shader, err = createAndCompileShader(string(fragShaderCode), gl.FRAGMENT_SHADER)
	if err != nil {
		return s, fmt.Errorf("%v: %v", vertexShaderPath, err)
	}
	s.attachShaderToProgram(shader)

	s.linkProgram()
	if err != nil {
		return s, fmt.Errorf("%v: %v", vertexShaderPath, err)
	}

	return s, nil
}
func (self *Shader) Delete() {
	gl.DeleteProgram(self.programID)
	self.programID = 0
}
func (self *Shader) Use() {
	gl.UseProgram(self.programID)
}
func (self *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(self.getUniformLocation(name), value)
}
func (self *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(self.getUniformLocation(name), value)
}
func (self *Shader) SetVec2(name string, value mgl32.Vec2) {
	gl.Uniform2fv(self.getUniformLocation(name), 1, &value[0])
}
func (self *Shader) SetVec3(name string, value mgl32.Vec3) {
	gl.Uniform3fv(self.getUniformLocation(name), 1, &value[0])
}
func (self *Shader) SetMat3(name string, value mgl32.Mat3) {
	gl.UniformMatrix3fv(self.getUniformLocation(name), 1, false, &value[0])
}
func (self *Shader) SetMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(self.getUniformLocation(name), 1, false, &value[0])
}

func (self *Shader) createProgram() {
	self.programID = gl.CreateProgram()
}
func (self *Shader) getUniformLocation(name string) int32 {
	name_cstr, free := gl.Strs(name)
	defer free()
	return gl.GetUniformLocation(self.programID, *name_cstr)
}
func createAndCompileShader(shaderCode string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	shaderCode = shaderCode + "\x00"
	csources, free := gl.Strs(shaderCode)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compile: %v", log)
	}

	return shader, nil
}
func (self *Shader) attachShaderToProgram(shader uint32) {
	gl.AttachShader(self.programID, shader)
	gl.DeleteShader(shader)
}
func (self *Shader) linkProgram() error {
	gl.LinkProgram(self.programID)

	var status int32
	gl.GetProgramiv(self.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(self.programID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(self.programID, logLength, nil, gl.Str(log))

		return fmt.Errorf("Failed to link program: %v", log)
	}
	return nil
}
