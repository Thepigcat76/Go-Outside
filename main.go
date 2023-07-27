package main

import (
	"go_outside/lib/logger"

	// "os"

	"embed"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//go:embed assets
var assets embed.FS

func main() {

	init_sdl()
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.SetResizable(true)

	imageData, err := assets.ReadFile("assets/textures/player.png")
	if err != nil {
		logger.Log("Failed to read file: "+err.Error(), logger.ERROR)
	}

	// Load the image into an SDL RWops
	rwops, err := sdl.RWFromMem(imageData)
	if err != nil {
		logger.Log("Failed to create RWops: "+sdl.GetError().Error(), logger.ERROR)
	}
	defer rwops.Close()

	surface_raw, err := img.LoadPNGRW(rwops) // Use 1 for freeing RWops after loading
	if err != nil {
		logger.Log("Failed to create Surface: "+err.Error(), logger.ERROR)
	}
	defer surface_raw.Free()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Create a texture from the surface
	texture, err := renderer.CreateTextureFromSurface(surface_raw)
	if err != nil {
		logger.Log("Failed to create Texture: "+err.Error(), logger.ERROR)
	}
	defer texture.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	// pixel := render_pixels(surface, [4]int32{255, 0, 255, 255})

	rect := sdl.FRect{X: 0, Y: 0, W: 80, H: 120}

	redColor := sdl.Color{R: 255, G: 255, B: 255, A: 255}

	keys := make([]bool, sdl.NUM_SCANCODES)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				keyEvent := event.(*sdl.KeyboardEvent)
				keyPressed := keyEvent.Keysym.Scancode
				// Check if player pressed escape to exit
				if keyPressed == sdl.SCANCODE_ESCAPE {
					logger.Log("Requested Exit", logger.WARNING)
					running = false
				}

				// Check for keybinds that are initialized
				// after the event loop
				if keyEvent.Type == sdl.KEYDOWN {
					// Set the corresponding key state to true when a key is pressed
					keys[keyEvent.Keysym.Scancode] = true
				} else if keyEvent.Type == sdl.KEYUP {
					// Set the corresponding key state to false when a key is released
					keys[keyEvent.Keysym.Scancode] = false
				}
			}
		}

		// Clear the renderer

		renderer.SetDrawColor(redColor.R, redColor.G, redColor.B, redColor.A)
		renderer.Clear()

		rectNormal := &sdl.Rect{
			X: int32(rect.X),
			Y: int32(rect.Y),
			W: int32(rect.W),
			H: int32(rect.H),
		}

		// Draw the image
		renderer.Copy(texture, nil, rectNormal)

		if keys[sdl.SCANCODE_W] {
			rect.Y -= 0.1
		}
		if keys[sdl.SCANCODE_S] {
			rect.Y += 0.1
		}
		if keys[sdl.SCANCODE_A] {
			rect.X -= 0.1
		}
		if keys[sdl.SCANCODE_D] {
			rect.X += 0.1
		}

		// Update the screen
		renderer.Present()
	}
}

func render_pixels(surface *sdl.Surface, rgba [4]int32) uint32 {
	color := sdl.Color{R: uint8(rgba[0]), G: uint8(rgba[1]), B: uint8(rgba[2]), A: uint8(rgba[3])} // purple
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	return pixel
}

func init_sdl() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	return err
}
