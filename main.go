package main

import (
	"game2d/core"
	"game2d/systems"
	"runtime"
	"time"
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

		window.BeginFrame()
		err := game.Update(elapsedTime.Seconds())
		if err != nil {
			panic(err)
		}
		// if game.IsOver {
		// 	break
		// }
		window.EndFrame()
	}
}
