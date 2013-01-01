package main

import (
	"fmt"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"github.com/banthar/gl"
	"math"
	"os"
)

var should_rotate = true

func quit_tutorial(code int) {
	sdl.Quit()
	os.Exit(code)
}

func handle_key_down(keysym *sdl.Keysym) {
	switch keysym.Sym {
	case sdl.K_ESCAPE:
		quit_tutorial(0)
	case sdl.K_SPACE:
		should_rotate = !should_rotate
	}
}

func process_events() {
outer:
	for {
		select {
		default:
			break outer
		case _event := <-sdl.Events:
			switch e := _event.(type) {
			case sdl.KeyboardEvent:
				handle_key_down(&e.Keysym)
			case sdl.QuitEvent:
				quit_tutorial(0)
			}
		}
	}
}

func draw_screen() {
	angle := float32(0.0)
	v := [][]float32{
		{-1.0, -1.0, 1.0},
		{1.0, -1.0, 1.0},
		{1.0, 1.0, 1.0},
		{-1.0, 1.0, 1.0},
		{-1.0, -1.0, -1.0},
		{1.0, -1.0, -1.0},
		{1.0, 1.0, -1.0},
		{-1.0, 1.0, -1.0}}
	red := []byte{255, 0, 0, 255}
	green := []byte{0, 255, 0, 255}
	blue := []byte{0, 0, 255, 255}
	white := []byte{255, 255, 255, 255}
	yellow := []byte{0, 255, 255, 255}
	black := []byte{0, 0, 0, 255}
	orange := []byte{255, 255, 0, 255}
	purple := []byte{255, 0, 255, 0}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -5.0)
	gl.Rotatef(angle, 0.0, 1.0, 0.0)
	if should_rotate {
		angle++
		if angle > 360.0 {
			angle = 0.0
		}
	}
	gl.Begin(gl.TRIANGLES)
	gl.Color4ubv(red)
	gl.Vertex3fv(v[0])
	gl.Color4ubv(green)
	gl.Vertex3fv(v[1])
	gl.Color4ubv(blue)
	gl.Vertex3fv(v[2])
	gl.Color4ubv(red)
	gl.Vertex3fv(v[0])
	gl.Color4ubv(blue)
	gl.Vertex3fv(v[2])
	gl.Color4ubv(white)
	gl.Vertex3fv(v[3])
	gl.Color4ubv(green)
	gl.Vertex3fv(v[1])
	gl.Color4ubv(black)
	gl.Vertex3fv(v[5])
	gl.Color4ubv(orange)
	gl.Vertex3fv(v[6])
	gl.Color4ubv(green)
	gl.Vertex3fv(v[1])
	gl.Color4ubv(orange)
	gl.Vertex3fv(v[6])
	gl.Color4ubv(blue)
	gl.Vertex3fv(v[2])
	gl.Color4ubv(black)
	gl.Vertex3fv(v[5])
	gl.Color4ubv(yellow)
	gl.Vertex3fv(v[4])
	gl.Color4ubv(purple)
	gl.Vertex3fv(v[7])
	gl.Color4ubv(black)
	gl.Vertex3fv(v[5])
	gl.Color4ubv(purple)
	gl.Vertex3fv(v[7])
	gl.Color4ubv(orange)
	gl.Vertex3fv(v[6])
	gl.Color4ubv(yellow)
	gl.Vertex3fv(v[4])
	gl.Color4ubv(red)
	gl.Vertex3fv(v[0])
	gl.Color4ubv(white)
	gl.Vertex3fv(v[3])
	gl.Color4ubv(yellow)
	gl.Vertex3fv(v[4])
	gl.Color4ubv(white)
	gl.Vertex3fv(v[3])
	gl.Color4ubv(purple)
	gl.Vertex3fv(v[7])
	gl.Color4ubv(white)
	gl.Vertex3fv(v[3])
	gl.Color4ubv(blue)
	gl.Vertex3fv(v[2])
	gl.Color4ubv(orange)
	gl.Vertex3fv(v[6])
	gl.Color4ubv(white)
	gl.Vertex3fv(v[3])
	gl.Color4ubv(orange)
	gl.Vertex3fv(v[6])
	gl.Color4ubv(purple)
	gl.Vertex3fv(v[7])
	gl.Color4ubv(green)
	gl.Vertex3fv(v[1])
	gl.Color4ubv(red)
	gl.Vertex3fv(v[0])
	gl.Color4ubv(yellow)
	gl.Vertex3fv(v[4])
	gl.Color4ubv(green)
	gl.Vertex3fv(v[1])
	gl.Color4ubv(yellow)
	gl.Vertex3fv(v[4])
	gl.Color4ubv(black)
	gl.Vertex3fv(v[5])
	gl.End()
	sdl.GL_SwapBuffers()
}

func setup_opengl(width, height int) {
	ratio := float64(width) / float64(height)
	gl.ShadeModel(gl.SMOOTH)
	gl.CullFace(gl.BACK)
	gl.FrontFace(gl.CCW)
	gl.Enable(gl.CULL_FACE)
	gl.ClearColor(0, 0, 0, 0)
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	gluPerspective := func(fovy, aspect, zNear, zFar float64) {
		top := math.Tan(fovy*0.5) * zNear
		bottom := -top
		left := aspect * bottom
		right := -left
		gl.Frustum(left, right, bottom, top, zNear, zFar)
	}
	gluPerspective(60.0, ratio, 1.0, 1024.0)
}

func main() {
	if sdl.Init(sdl.INIT_VIDEO) < 0 {
		fmt.Printf(
			"Video initialization failed: %s\n",
			sdl.GetError())
		quit_tutorial(1)
	}
	info := sdl.GetVideoInfo()
	if info == nil {
		fmt.Printf(
			"Video query failed: %s\n",
			sdl.GetError())
		quit_tutorial(1)
	}
	width := 640
	height := 480
	bpp := int(info.Vfmt.BitsPerPixel)
	sdl.GL_SetAttribute(sdl.GL_RED_SIZE, 5)
	sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, 5)
	sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, 5)
	sdl.GL_SetAttribute(sdl.GL_DEPTH_SIZE, 16)
	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	flags := uint32(sdl.OPENGL | sdl.FULLSCREEN)
	if sdl.SetVideoMode(width, height, bpp, flags) == nil {
		fmt.Printf("Video mode set failed: %s\n", sdl.GetError())
		quit_tutorial(1)
	}
	setup_opengl(width, height)
	for {
		process_events()
		draw_screen()
	}
	os.Exit(0)
}
