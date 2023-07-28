package main

import (
	"go_outside/lib/logger"
	"go_outside/lib"

	// "os"

	"embed"

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


	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger.Log("Failed to create Renderer: "+err.Error(), logger.ERROR)
	}
	defer renderer.Destroy()

	texture, err := lib.Load_image("assets/textures/player.png", renderer, assets)
	defer texture.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		logger.Log("Failed to create Surface: "+sdl.GetError().Error(), logger.ERROR)
	}
	surface.FillRect(nil, 0)

	// pixel := render_pixels(surface, [4]int32{255, 0, 255, 255})

	rect := sdl.FRect{X: 0, Y: 0, W: 80, H: 120}

	rect2 := sdl.Rect{X: 0, Y: 0, W: 80, H: 120}

	// redColor := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	blueColor := sdl.Color{R: 0, G: 0, B: 255, A: 255}

	keys := make([]bool, sdl.NUM_SCANCODES)

	cursorInside := false
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
				winEvent := event.(*sdl.WindowEvent)
				if winEvent.Event == sdl.WINDOWEVENT_ENTER {
					cursorInside = true
					logger.Log("Cursor entered the window.", logger.INFO)
				} else if winEvent.Event == sdl.WINDOWEVENT_LEAVE {
					cursorInside = false
					logger.Log("Cursor left the window.", logger.INFO)
				}
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
			case *sdl.MouseButtonEvent:
				// Check if it's a mouse button down event
				mouseEvent := event.(*sdl.MouseButtonEvent)
				if mouseEvent.Type == sdl.MOUSEBUTTONDOWN {
					// Left mouse button pressed
					if mouseEvent.Button == sdl.BUTTON_LEFT {
					}
					// Right mouse button pressed
					if mouseEvent.Button == sdl.BUTTON_RIGHT {
					}
				}
			}
			
		}

		mouseX, mouseY, _ := sdl.GetMouseState()
		leftClick, _ := checkMouseClick()


		// Clear the renderer

		renderer.SetDrawColor(255, 0, 0, 255)

		if checkCollision(&rect2, mouseX, mouseY) && cursorInside {
			renderer.SetDrawColor(0, 0, 100, 255)
			if leftClick {
				renderer.SetDrawColor(blueColor.R, blueColor.G, blueColor.B, blueColor.A)
			}
		}

		renderer.Clear()
		renderer.SetDrawColor(100, 100, 100, 255)
		renderer.FillRect(&rect2)

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

func checkCollision(rect *sdl.Rect, mouseX, mouseY int32) bool {
	// Check if the mouse position is within the bounding rectangle
	return mouseX >= rect.X && mouseX <= rect.X+rect.W && mouseY >= rect.Y && mouseY <= rect.Y+rect.H
}

func checkMouseClick() (leftClick, rightClick bool) {
	// Get the current state of the mouse buttons
	_, _, mouseState := sdl.GetMouseState()

	// Check if the left mouse button is clicked
	leftClick = mouseState&sdl.ButtonLMask() != 0

	// Check if the right mouse button is clicked
	rightClick = mouseState&sdl.ButtonRMask() != 0

	return leftClick, rightClick
}