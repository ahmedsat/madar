package main

import (
	"fmt"
	"image/color"
	"os"
	"runtime"

	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
	"github.com/go-gl/gl/v4.6-core/gl"
)

func init() {
	runtime.LockOSThread()
}

var vertices = []noor.Vertex{
	{Position: [3]float32{0, 0.5, 0}, Color: [3]float32{1, 0, 0}, UV: [2]float32{0.5, 1}},
	{Position: [3]float32{0.5, -0.5, 0}, Color: [3]float32{0, 1, 0}, UV: [2]float32{1, 0}},
	{Position: [3]float32{-0.5, -0.5, 0}, Color: [3]float32{0, 0, 1}, UV: [2]float32{0, 0}},
}

func main() {

	n := noor.New(800, 600, "Hello, Shader!", color.Black).UnwrapOrPanic()
	defer n.Close()
	n.SetBackground(color.RGBA{R: 0x20, G: 0x30, B: 0x30, A: 0xff})

	// ? notice that if loading our shader fails `noor` will us its default shaders
	shader := noor.CreateShaderProgramFromFiles(
		"examples/shader.vert",
		"examples/shader.frag",
	).UnwrapOrPanic()
	defer shader.Delete()
	n.SetShader(shader)

	n.Shader.Activate()
	location := getUniformLocation(n.Shader, "matrix")

	r := float32(0)

	mesh := noor.NewMesh(vertices, nil, noor.DrawTriangles)

	for !n.ShouldClose() {

		r += 1
		if r > 360 {
			r = 0
		}

		var mat madar.Matrix
		mat = madar.IdentityMatrix4X4()
		mat = madar.RotationMatrix(r, r, r)
		gl.UniformMatrix4fv(location, 1, false, mat.GL())

		mesh.Draw()
	}
}

func getUniformLocation(sh noor.Shader, name string) int32 {

	nameCStr := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(uint32(sh), nameCStr)
	if location == -1 {
		fmt.Fprintf(os.Stderr, "Uniform location not found: %s\n", name)
	}
	return location
}
