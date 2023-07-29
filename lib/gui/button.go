package gui

import (
	"embed"

	"go_outside/lib"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Name string
	Clicked bool
	texture_std *sdl.Texture
	texture_sel *sdl.Texture
	button_rect *sdl.Rect
}

func Create_button(name string, texturePath string, renderer *sdl.Renderer, assets embed.FS, X, Y, W, H int32) Button {
	texture_std := lib.Load_image(texturePath, renderer, assets)
	texture_sel := lib.Load_image(texturePath + "_selected", renderer, assets)
	button_rect := sdl.Rect{X: X, Y: Y, W: W, H: H}

	return Button{Name: name, Clicked: false, texture_std: texture_std, texture_sel: texture_sel, button_rect: &button_rect}
}

func (b *Button) Draw_button(renderer *sdl.Renderer) {
	mouse_x, mouse_y, _ := sdl.GetMouseState()
	left_click, _ := lib.Mouse_clicked()
	if lib.Collide_point(b.button_rect, mouse_x, mouse_y) {
		renderer.Copy(b.texture_sel, nil ,b.button_rect)
		if left_click {
			b.Clicked = true
		} else {
			b.Clicked = false
		}
	} else {
		renderer.Copy(b.texture_std, nil ,b.button_rect)
	}
}
