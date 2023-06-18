package core

import (
	"game2d/config"

	"github.com/go-gl/mathgl/mgl32"
)

type CameraMovement int

const (
	CAMERA_RIGHT CameraMovement = iota
	CAMERA_UP
	CAMERA_LEFT
	CAMERA_DOWN
)

type Camera struct {
	position   mgl32.Vec3
	front      mgl32.Vec3
	up         mgl32.Vec3
	right      mgl32.Vec3
	screenSize mgl32.Vec2
}

func CreateCamera() Camera {
	var camera Camera
	camera.position = mgl32.Vec3{0.0, float32(config.C.CameraYOffset), 0.0}
	camera.front = mgl32.Vec3{0.0, 0.0, -1.0}
	camera.up = mgl32.Vec3{0.0, 1.0, 0.0}
	camera.right = mgl32.Vec3{1.0, 0.0, 0.0}
	camera.screenSize = mgl32.Vec2{float32(config.C.ScreenWidth), float32(config.C.ScreenHeight)}
	return camera
}

func (cam *Camera) GetViewMatrix() mgl32.Mat4 {
	dir := cam.position.Add(cam.front)
	return mgl32.LookAt(cam.position[0], cam.position[1], cam.position[2], dir[0], dir[1], dir[2], cam.up[0], cam.up[1], cam.up[2])
}
func (cam *Camera) GetProjectionMatrix() mgl32.Mat4 {
	return mgl32.Ortho(-cam.screenSize.X()/2, cam.screenSize.X()/2, -cam.screenSize.Y()/2, cam.screenSize.Y()/2, 0, 1000)
}
