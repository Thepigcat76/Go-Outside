package gui

import (
	"embed"

	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Name        string
	Clicked     bool
	Visible     bool
	texture_std *sdl.Texture
	texture_sel *sdl.Texture
	Button_rect *sdl.Rect
	X           int32
	Y           int32
}

func Create_button(name string, texturePath string, renderer *sdl.Renderer, assets embed.FS, X, Y, W, H int32, visible bool) Button {
	texture_std := util.LoadImage(texturePath, renderer, assets, 1, true)
	texture_sel := util.LoadImage(texturePath+"_selected", renderer, assets, 1, true)
	button_rect := sdl.Rect{X: X, Y: Y, W: W, H: H}

	if !visible {
		return Button{Name: name, Clicked: false, Visible: false, texture_std: texture_std.Texture, texture_sel: texture_sel.Texture, Button_rect: &button_rect}
	}

	return Button{Name: name, Clicked: false, Visible: true, texture_std: texture_std.Texture, texture_sel: texture_sel.Texture, Button_rect: &button_rect}
}

func (b *Button) Draw_button(renderer *sdl.Renderer) {
	b.Button_rect.X, b.Button_rect.Y = b.X, b.Y
	mouse_x, mouse_y, _ := sdl.GetMouseState()
	left_click, _ := util.Mouse_clicked()
	if b.Visible {
		if util.Collide_point(b.Button_rect, mouse_x, mouse_y) {
			renderer.Copy(b.texture_sel, nil, b.Button_rect)
			if left_click {
				b.Clicked = true
			} else {
				b.Clicked = false
			}
		} else {
			renderer.Copy(b.texture_std, nil, b.Button_rect)
		}
	}
}

func (b *Button) Collide_mouse() bool {
	mouse_x, mouse_y, _ := sdl.GetMouseState()
	if b.Visible {
		if util.Collide_point(b.Button_rect, mouse_x, mouse_y) {
			return true
		}
	}
	return false
}
