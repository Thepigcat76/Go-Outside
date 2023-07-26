package main

import (
	// "go_outside/lib"
	// "fmt"
	"go_outside/lib"
	"go_outside/lib/logger"
	// "os"

	// "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	init_sdl()
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Load the image
	imagePath := "assets/textures/test.png" // Replace this with the actual path to your image
	texture, err := lib.Load_image(imagePath, renderer)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	pixel := render_pixels(surface, [4]int32{255, 0, 255, 255})

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	surface.FillRect(&rect, pixel)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// Clear the renderer
		renderer.Clear()

		// Draw the image
		renderer.Copy(texture, nil, nil)

		// Update the screen
		renderer.Present()

		logger.Log("amogus", logger.SUCCESS)
	}
}


func render_pixels(surface *sdl.Surface, rgba [4]int32) uint32 {
	color := sdl.Color{R: uint8(rgba[0]), G: uint8(rgba[1]), B: uint8(rgba[2]), A: uint8(rgba[3])} // purple
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	return pixel
}

func init_sdl() error {
	err := sdl.Init(sdl.INIT_EVERYTHING);
	if err != nil {
		panic(err)
	}
	return err
}