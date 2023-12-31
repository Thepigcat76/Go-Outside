package main

import (
	"go_outside/lib/gui"
	"go_outside/lib/inventory"
	"go_outside/lib/item"
	"go_outside/lib/logger"
	"go_outside/lib/player"
	"go_outside/lib/util"

	"embed"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//go:embed assets
var assets embed.FS

const escapeCooldown = 500

func main() {
	font_path := "assets/fonts/FFFFORWA.TTF"

	util.Init_sdl()
	defer sdl.Quit()
	defer ttf.Quit()

	window := util.Create_window(800, 600, true)
	defer window.Destroy()

	renderer := util.Create_renderer(window)
	defer renderer.Destroy()

	font := util.LoadFont(font_path, 12, "Amogus", &sdl.Color{R: 255, G: 255, B: 255}, renderer, assets)
	defer font.Surface.Free()
	defer font.Texture.Destroy()
	defer font.Font.Close()

	world, err := sdl.CreateRGBSurface(0, 800, 600, 32, 0, 0, 0, 0)
	if err != nil {
		// Handle error
		panic(err)
	}
	defer world.Free()

	world_texture, err := renderer.CreateTextureFromSurface(world)
	if err != nil {
		// Handle error
		panic(err)
	}

	items := item.InitItems(renderer, assets, font_path)

	test_sword := items.New("test_sword", item.Settings{Rarity: item.COMMON, Durability: 100, Tooltip: "amogus"})
	test_sword_1 := items.New("copper_gear", item.Settings{Rarity: item.COMMON, Durability: 100, Tooltip: "amogus"})
	println(test_sword_1.Texture.X, test_sword_1.Texture.Y)

	player := player.Create([4]string{"assets/textures/player"}, renderer, assets, 5.0, 200, 200)

	keys_pressed := make([]bool, sdl.NUM_SCANCODES)

	quit_button := gui.Create_button("quit_button", "assets/textures/quit_button", renderer, assets, 600, 200, 100, 100, false)

	options_button := gui.Create_button("options_button", "assets/textures/options_button", renderer, assets, 400, 200, 100, 100, false)

	continue_button := gui.Create_button("continue_button", "assets/textures/continue_button", renderer, assets, 200, 200, 100, 100, false)

	inventory := inventory.Init_inventory(renderer, assets)

	test_texture := util.LoadImage("assets/textures/button_template_selected", renderer, assets, 4, true)

	running := true
	var slot int32 = 0
	test_rect := sdl.Rect{X: 100, Y: 100, W: 1000, H: 1000}

	// destRect := sdl.FRect{X: 100, Y: 200, W: float32(world.W), H: float32(world.H)}

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
					keys_pressed[key_event.Keysym.Scancode] = true
				} else if key_event.Type == sdl.KEYUP {
					// Set the corresponding key state to false when a key is released
					keys_pressed[key_event.Keysym.Scancode] = false

					// Check if escape key was pressed
					if key_pressed == sdl.SCANCODE_ESCAPE {
						quit_button.Visible = !quit_button.Visible
						options_button.Visible = !options_button.Visible
						continue_button.Visible = !continue_button.Visible
					}
					if key_pressed == sdl.SCANCODE_P {
						inventory.Set_item(*test_sword, slot)
						slot += 1
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

		if keys_pressed[sdl.SCANCODE_W] {
			test_rect.Y -= 1
		}
		if keys_pressed[sdl.SCANCODE_S] {
			test_rect.Y += 1
		}
		if keys_pressed[sdl.SCANCODE_A] {
			test_rect.X -= 1
		}
		if keys_pressed[sdl.SCANCODE_D] {
			test_rect.X += 1
		}

		renderer.Copy(world_texture, nil, &test_rect)

		renderer.SetRenderTarget(world_texture)

		for y := 0; y < 30; y++ {
			for x := 0; x < 2; x++ {
				renderer.Copy(test_texture.Texture, nil, &test_rect)
			}
		}

		test_sword_1_Y := float32(600.0)

		world.FillRect(nil, sdl.MapRGB(surface.Format, 0, 0, 0))

		test_sword.Draw_single(&test_sword.Texture.X, &test_sword.Texture.Y)
		
		test_sword_1.Draw_single(&test_sword.Texture.X, &test_sword_1_Y)
		println(test_sword_1.Texture.X, test_sword_1.Texture.Y)

		player.Draw()

		inventory.Draw(items, 0, 100)

		item.Draw(items)

		font.Draw(200, 100)

		quit_button.Draw_button(renderer)
		options_button.Draw_button(renderer)
		continue_button.Draw_button(renderer)

		// TODO: put this in method
		continue_button.X = surface.W / 5
		options_button.X = surface.W/5 + surface.W/4
		quit_button.X = surface.W/5 + surface.W/2

		continue_button.Y = surface.H / 3
		options_button.Y = surface.H / 3
		quit_button.Y = surface.H / 3

		// Update the screen
		renderer.Present()
	}
}
