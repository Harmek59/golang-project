package main

import (
	"game2d/core"
	"game2d/systems"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var window core.Window

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func main() {
	var err error
	window, err = core.CreateWindow()
	defer window.Delete()
	if err != nil {
		panic(err)
	}

	lastTime := time.Now()
	game := systems.NewGame()
	for !window.ShouldClose() {
		currentTime := time.Now()
		elapsedTime := currentTime.Sub(lastTime)
		lastTime = currentTime
		if glfw.GetCurrentContext().GetKey(glfw.KeyEscape) == glfw.Press {
            break
		}

		window.BeginFrame()
		err := game.Update(elapsedTime.Seconds())
		if err != nil {
			panic(err)
		}
		window.EndFrame()
	}
}
