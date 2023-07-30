package main

import (
	"go_outside/lib/gui"
	"go_outside/lib/logger"
	// "go_outside/lib/player"
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

	texture := util.Load_image("assets/textures/player", renderer, assets, 10)

	// player := player.Create_player()

	rect := sdl.FRect{X: 0, Y: 0, W: 80, H: 120}

	keys := make([]bool, sdl.NUM_SCANCODES)

	show_escape_menu := false

	quit_button := gui.Create_button("quit_button", "assets/textures/quit_button", renderer, assets, 600, 200, 100, 100, false)

	options_button := gui.Create_button("options_button", "assets/textures/options_button", renderer, assets, 400, 200, 100, 100, false)

	continue_button := gui.Create_button("continue_button", "assets/textures/continue_button", renderer, assets, 200, 200, 100, 100, false)

	buttons := [3]gui.Button{quit_button, options_button, continue_button}

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
						show_escape_menu = !show_escape_menu
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
			}

		}

		surface := util.Create_surface_from_window(window)
		
		// Clear the renderer
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()

		// Draw the image
		texture.Draw_image(renderer, 0, 0)

		if show_escape_menu {
			quit_button.Visible = true
			options_button.Visible = true
			continue_button.Visible = true
		} else {
			quit_button.Visible = false
			options_button.Visible = false
			continue_button.Visible = false
		}

		if quit_button.Clicked {
			logger.Log("Requested Exit", logger.WARNING)
			running = false
			break
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

		handle_button_pos(buttons[:], surface)

		quit_button.Draw_button(renderer)
		options_button.Draw_button(renderer)
		continue_button.Draw_button(renderer)

		// Update the screen
		renderer.Present()
	}
}

func handle_button_pos(buttons []gui.Button, surface *sdl.Surface) {
    for i := 0; i < len(buttons); i++ {
		buttons[i].Button_rect.X = surface.W / 3
		buttons[i].Button_rect.X += int32(i * 200)
    }
	println(surface.W)
}
