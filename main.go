package main

import (
	"go_outside/lib"
	"go_outside/lib/logger"
	"go_outside/lib/gui"

	"embed"

	"github.com/veandco/go-sdl2/sdl"
)

//go:embed assets
var assets embed.FS

func main() {
	lib.Init_sdl()
	defer sdl.Quit()

	window := lib.Create_window(800, 600, true)
	defer window.Destroy()

	renderer := lib.Create_renderer(window)
	defer renderer.Destroy()

	texture := lib.Load_image("assets/textures/player", renderer, assets)
	defer texture.Destroy()

	rect := sdl.FRect{X: 0, Y: 0, W: 80, H: 120}

	keys := make([]bool, sdl.NUM_SCANCODES)

	test_button := gui.Create_button("test_button", "assets/textures/test_button", renderer, assets, 200, 200, 100, 100)
	test_button.Visible = false

	running := true
	logger.Log("Started successfully", logger.SUCCESS)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				logger.Log("Requested Exit", logger.WARNING)
				break
			case *sdl.KeyboardEvent:
				key_event := event.(*sdl.KeyboardEvent)
				key_pressed := key_event.Keysym.Scancode
				// Check if player pressed escape to exit
				if key_pressed == sdl.SCANCODE_COMMA {
					logger.Log("Requested Exit", logger.WARNING)
					running = false
				}

				// Check for keybinds that are initialized
				// after the event loop
				if key_event.Type == sdl.KEYDOWN {
					// Set the corresponding key state to true when a key is pressed
					keys[key_event.Keysym.Scancode] = true
				} else if key_event.Type == sdl.KEYUP {
					// Set the corresponding key state to false when a key is released
					keys[key_event.Keysym.Scancode] = false
				}
			case *sdl.MouseButtonEvent:
				// Check if it's a mouse button down event
				mouse_event := event.(*sdl.MouseButtonEvent)
				if mouse_event.Type == sdl.MOUSEBUTTONDOWN {
					// Left mouse button pressed
					if mouse_event.Button == sdl.BUTTON_LEFT {
					}
					// Right mouse button pressed
					if mouse_event.Button == sdl.BUTTON_RIGHT {
					}
				}
			}
			
		}

		// Clear the renderer
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()

		rect_normal := lib.Convert_frect_to_rect(&rect)

		// Draw the image
		renderer.Copy(texture, nil, &rect_normal)

		if keys[sdl.SCANCODE_ESCAPE] {
			test_button.Visible = true
		}

		if test_button.Clicked {
			rect.X = 500
		}

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
		test_button.Draw_button(renderer)

		// Update the screen
		renderer.Present()
	}
}
