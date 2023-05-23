package core

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	WIDTH       = 500
	HEIGHT      = 500
	WINDOW_NAME = "Game2D"
)

type Window struct {
	window *glfw.Window
}

func CreateWindow() (Window, error) {
	var window Window
	// init glfw
	err := glfw.Init()
	if err != nil {
		return window, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window.window, err = glfw.CreateWindow(WIDTH, HEIGHT, WINDOW_NAME, nil, nil)
	if err != nil {
		return window, err
	}
	window.window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		return window, err
	}
	gl.Viewport(0, 0, WIDTH, HEIGHT)

	// Vsync
	glfw.SwapInterval(1)

	return window, nil
}
func (self *Window) Delete() {
	self.window.Destroy()
	glfw.Terminate()
}
func (self *Window) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
func (self *Window) EndFrame() {
	glfw.PollEvents()
	self.window.SwapBuffers()
}
func (self *Window) ShouldClose() bool {
	return self.window.ShouldClose()
}
func (self *Window) GetWindow() *glfw.Window {
    return self.window
}
