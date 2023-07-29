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

	// surface := create_surface_from_window(window)

	// pixel := render_pixels(surface, [4]int32{255, 0, 255, 255})

	rect := sdl.FRect{X: 0, Y: 0, W: 80, H: 120}

	rect2 := sdl.Rect{X: 0, Y: 0, W: 80, H: 120}

	// redColor := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	blue_color := sdl.Color{R: 0, G: 0, B: 255, A: 255}

	keys := make([]bool, sdl.NUM_SCANCODES)

	test_button := gui.Create_button("test_button", "assets/textures/test_button", renderer, assets, 200, 200, 100, 100)

	cursor_inside := false
	running := true
	logger.Log("Started successfully", logger.SUCCESS)
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				logger.Log("Requested Exit", logger.WARNING)
				break
			case *sdl.WindowEvent:
				// Check if it's a window event related to mouse motion
				win_event := event.(*sdl.WindowEvent)
				if win_event.Event == sdl.WINDOWEVENT_ENTER {
					cursor_inside = true
					logger.Log("Cursor entered the window.", logger.INFO)
				} else if win_event.Event == sdl.WINDOWEVENT_LEAVE {
					cursor_inside = false
					logger.Log("Cursor left the window.", logger.INFO)
				}
			case *sdl.KeyboardEvent:
				key_event := event.(*sdl.KeyboardEvent)
				key_pressed := key_event.Keysym.Scancode
				// Check if player pressed escape to exit
				if key_pressed == sdl.SCANCODE_ESCAPE {
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

		mouse_x, mouse_y, _ := sdl.GetMouseState()
		left_click, _ := lib.Mouse_clicked()

		// Clear the renderer

		renderer.SetDrawColor(255, 0, 0, 255)

		if lib.Collide_point(&rect2, mouse_x, mouse_y) && cursor_inside {
			renderer.SetDrawColor(0, 0, 100, 255)
			if left_click {
				renderer.SetDrawColor(blue_color.R, blue_color.G, blue_color.B, blue_color.A)
			}
		}

		renderer.Clear()
		renderer.SetDrawColor(100, 100, 100, 255)
		renderer.FillRect(&rect2)

		rect_normal := lib.Convert_frect_to_rect(&rect)

		// Draw the image
		renderer.Copy(texture, nil, &rect_normal)

		test_button.Draw_button(renderer)

		if test_button.Clicked {
			logger.Log("clicked " + test_button.Name, logger.INFO)
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

		// Update the screen
		renderer.Present()
	}
}
