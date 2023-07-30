package util

import (
	"go_outside/lib/logger"
	"github.com/veandco/go-sdl2/sdl"
)

func Render_pixels(surface *sdl.Surface, rgba [4]int32) uint32 {
	color := sdl.Color{R: uint8(rgba[0]), G: uint8(rgba[1]), B: uint8(rgba[2]), A: uint8(rgba[3])} // purple
	pixel := sdl.MapRGBA(surface.Format, color.R, color.G, color.B, color.A)
	return pixel
}

func Init_sdl() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	return err
}

func Collide_point(rect *sdl.Rect, X, Y int32) bool {
	// Check if the mouse position is within the bounding rectangle
	return X >= rect.X && X <= rect.X+rect.W && Y >= rect.Y && Y <= rect.Y+rect.H
}

func Mouse_clicked() (leftClick, rightClick bool) {
	// Get the current state of the mouse buttons
	_, _, mouseState := sdl.GetMouseState()

	// Check if the left mouse button is clicked
	leftClick = mouseState&sdl.ButtonLMask() != 0

	// Check if the right mouse button is clicked
	rightClick = mouseState&sdl.ButtonRMask() != 0

	return leftClick, rightClick
}

func Create_window(width int32, height int32, allowResize bool) *sdl.Window {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	if allowResize {
		window.SetResizable(true)
	} else {
		window.SetResizable(false)
	}
	return window
}

func Create_renderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		logger.Log("Failed to create Renderer: "+err.Error(), logger.ERROR)
	}
	return renderer
}

func Create_surface_from_window(window *sdl.Window) *sdl.Surface {
	surface, err := window.GetSurface()
	if err != nil {
		logger.Log("Failed to create Surface: "+sdl.GetError().Error(), logger.ERROR)
	}
	return surface
}

func Convert_frect_to_rect(frect *sdl.FRect) sdl.Rect {
	rect := &sdl.Rect{
		X: int32(frect.X),
		Y: int32(frect.Y),
		W: int32(frect.W),
		H: int32(frect.H),
	}
	return *rect
}

type Timer struct {
	Time       int32
	Reset_time int32
	Reset_bool bool
	running    bool
}

func Create_timer() Timer {
	return Timer{Time: 0, Reset_time: 0, Reset_bool: false, running: false}
}

func (t *Timer) Start() {
	t.running = true
}

// run has to be called for every timer
// regardless whether it is actually running or not
func (t *Timer) Run() {
	if t.running == true {
		t.Time++
	}
}

func (t *Timer) Stop() {
	t.running = false
}

func (t *Timer) Reset() {
	t.Time = 0
	t.Reset_time++
	t.Reset_bool = true
}

func (t *Timer) Stop_and_reset() {
	t.running = false
	t.Time = 0
	t.Reset_time++
	t.Reset_bool = true
}
