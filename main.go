package main

import (
	"fmt"
	"runtime"
	"strconv"

	"game2D/core"
	"game2D/ogl"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/rivo/tview"
)

var mainShader ogl.Shader
var vertices ogl.Vertices
var window core.Window
var app *tview.Application

var color mgl32.Vec3 = mgl32.Vec3{1.0, 0.5, 0.4}

const (
	width  = 500
	height = 500
)

var (
	triangles = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,

		0.5, 0.5, 0,
		0.5, -0.5, 0,
		-0.5, 0.5, 0,
	}
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	runtime.LockOSThread()
}

func run_tview() {
	float_accept := func(val string, r rune) bool {
		_, e := strconv.ParseFloat(val, 32)
		if e != nil {
			return false
		}
		return true
	}
	float_changed := func(val string, valToChange *float32) {
		f, e := strconv.ParseFloat(val, 32)
		if e == nil {
			*valToChange = float32(f)
		}
	}
	floatToStr := func(f float32) string {
		return fmt.Sprintf("%f", f)

	}
	app = tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Color R", floatToStr(color[0]), 20, float_accept, func(val string) { float_changed(val, &color[0]) }).
		AddInputField("Color G", floatToStr(color[1]), 20, float_accept, func(val string) { float_changed(val, &color[1]) }).
		AddInputField("Color B", floatToStr(color[2]), 20, float_accept, func(val string) { float_changed(val, &color[2]) }).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Change color").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func main() {
	var err error

	window, err = core.CreateWindow()
	defer window.Delete()
	if err != nil {
		panic(err)
	}

	vector := mgl32.Vec3{1.0, 0.0, 0.0}
	fmt.Println(vector)
	fmt.Println(color[0])

	mainShader, err = ogl.CreateShader("shaders/mainShader.vert", "shaders/mainShader.frag")
	defer mainShader.Delete()
	if err != nil {
		panic(err)
	}
	buff := ogl.CreateBuffer(gl.ARRAY_BUFFER, triangles)
	defer buff.Delete()
	buff.Bind(gl.ARRAY_BUFFER)

	vertices = ogl.CreateVertices(triangles, []uint32{3})
	defer vertices.Delete()

	go run_tview()

	for !window.ShouldClose() {
		draw()
	}
    app.Stop()
}

func draw() {
	window.BeginFrame()

	mainShader.Use()
	mainShader.SetVec3("color", color)
	vertices.Bind()
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangles)/3))

	window.EndFrame()
}
