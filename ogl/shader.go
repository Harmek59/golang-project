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
		return s, fmt.Errorf("Failed to load file: %v, err: %v", fragmentShaderPath, err)
	}
	shader, err = createAndCompileShader(string(fragShaderCode), gl.FRAGMENT_SHADER)
	if err != nil {
		return s, fmt.Errorf("%v: %v", fragmentShaderPath, err)
	}
	s.attachShaderToProgram(shader)

	s.linkProgram()
	if err != nil {
		return s, fmt.Errorf("%v: %v", vertexShaderPath, err)
	}

	return s, nil
}
func (s *Shader) Delete() {
	gl.DeleteProgram(s.programID)
	s.programID = 0
}
func (s *Shader) Use() {
	gl.UseProgram(s.programID)
}
func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(s.getUniformLocation(name), value)
}
func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(s.getUniformLocation(name), value)
}
func (s *Shader) SetVec2(name string, value mgl32.Vec2) {
	gl.Uniform2fv(s.getUniformLocation(name), 1, &value[0])
}
func (s *Shader) SetVec3(name string, value mgl32.Vec3) {
	gl.Uniform3fv(s.getUniformLocation(name), 1, &value[0])
}
func (s *Shader) SetMat3(name string, value mgl32.Mat3) {
	gl.UniformMatrix3fv(s.getUniformLocation(name), 1, false, &value[0])
}
func (s *Shader) SetMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(s.getUniformLocation(name), 1, false, &value[0])
}

func (s *Shader) createProgram() {
	s.programID = gl.CreateProgram()
}
func (s *Shader) getUniformLocation(name string) int32 {
	name_cstr, free := gl.Strs(name)
	defer free()
	return gl.GetUniformLocation(s.programID, *name_cstr)
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
func (s *Shader) attachShaderToProgram(shader uint32) {
	gl.AttachShader(s.programID, shader)
	gl.DeleteShader(shader)
}
func (s *Shader) linkProgram() error {
	gl.LinkProgram(s.programID)

	var status int32
	gl.GetProgramiv(s.programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(s.programID, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(s.programID, logLength, nil, gl.Str(log))

		return fmt.Errorf("Failed to link program: %v", log)
	}
	return nil
}
