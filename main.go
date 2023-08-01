package main

import (
	"go_outside/lib/gui"
	"go_outside/lib/inventory"
	"go_outside/lib/logger"
	"go_outside/lib/player"
	"go_outside/lib/util"

	"embed"

	"github.com/veandco/go-sdl2/sdl"
)

//go:embed assets
var assets embed.FS

const escapeCooldown = 500

func main() {

	util.Init_sdl()
	defer sdl.Quit()

	window := util.Create_window(800, 600, true)
	defer window.Destroy()

	renderer := util.Create_renderer(window)
	defer renderer.Destroy()

	player := player.Create([4]string{"assets/textures/player"}, renderer, assets, 5.0, 200, 200)

	keys := make([]bool, sdl.NUM_SCANCODES)

	quit_button := gui.Create_button("quit_button", "assets/textures/quit_button", renderer, assets, 600, 200, 100, 100, false)

	options_button := gui.Create_button("options_button", "assets/textures/options_button", renderer, assets, 400, 200, 100, 100, false)

	continue_button := gui.Create_button("continue_button", "assets/textures/continue_button", renderer, assets, 200, 200, 100, 100, false)

	texture := util.Load_image("assets/textures/infinity_sword", renderer, assets, 5.0)

	inventory := inventory.Init_inventory(renderer, assets)

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

				// Check for keybinds that are initialized
				// after the event loop
				if key_event.Type == sdl.KEYDOWN {
					// Set the corresponding key state to true when a key is pressed
					keys[key_event.Keysym.Scancode] = true
				} else if key_event.Type == sdl.KEYUP {
					// Set the corresponding key state to false when a key is released
					keys[key_event.Keysym.Scancode] = false

					// Check if escape key was pressed
					if key_pressed == sdl.SCANCODE_ESCAPE {
						quit_button.Visible = !quit_button.Visible
						options_button.Visible = !options_button.Visible
						continue_button.Visible = !continue_button.Visible
					}
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
				if mouse_event.Type == sdl.MOUSEBUTTONUP {
					if mouse_event.Button == sdl.BUTTON_LEFT {
						if continue_button.Collide_mouse() {
							quit_button.Visible = false
							options_button.Visible = false
							continue_button.Visible = false
						}
					}
				}
			}

		}

		surface := util.Create_surface_from_window(window)
		
		// Clear the renderer
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()


		if quit_button.Clicked {
			logger.Log("Requested Exit", logger.WARNING)
			running = false
			break
		}

		player.Draw()
		texture.Draw_image()
		texture.X = 266

		if keys[sdl.SCANCODE_W] {
			player.Y -= 0.1
		}
		if keys[sdl.SCANCODE_S] {
			player.Y += 0.1
		}
		if keys[sdl.SCANCODE_A] {
			player.X -= 0.1
		}
		if keys[sdl.SCANCODE_D] {
			player.X += 0.1
		}

		quit_button.Draw_button(renderer)
		options_button.Draw_button(renderer)
		continue_button.Draw_button(renderer)

		inventory.Draw_Inventory(0, 100)

		continue_button.X = surface.W / 5
		options_button.X = surface.W / 5 + surface.W / 4
		quit_button.X = surface.W / 5 + surface.W / 2

		continue_button.Y = surface.H / 3
		options_button.Y = surface.H / 3
		quit_button.Y = surface.H / 3

		// Update the screen
		renderer.Present()
	}
}