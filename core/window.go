package core

import (
	"game2d/config"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	WindowName = "Game2D"
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

	window.window, err = glfw.CreateWindow(int(config.C.ScreenWidth), int(config.C.ScreenHeight), WindowName, nil, nil)
	if err != nil {
		return window, err
	}
	window.window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		return window, err
	}
	gl.Viewport(0, 0, int32(config.C.ScreenWidth), int32(config.C.ScreenHeight))

	// Vsync
	glfw.SwapInterval(1)
    gl.Enable(gl.BLEND)
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);

	return window, nil
}
func (window *Window) Delete() {
	window.window.Destroy()
	glfw.Terminate()
}
func (window *Window) BeginFrame() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
func (window *Window) EndFrame() {
	glfw.PollEvents()
	window.window.SwapBuffers()
}
func (window *Window) ShouldClose() bool {
	return window.window.ShouldClose()
}
func (window *Window) GetWindow() *glfw.Window {
	return window.window
}
