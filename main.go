package main

import (
	// "go_outside/lib"
	"go_outside/lib/logger"

	// "github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// initialize sdl2
	init_sdl()
	// Quit at the end
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	pixel := render_pixels(surface, [4]int32{255, 0, 255, 255})

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
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