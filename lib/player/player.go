package player

import (
	"embed"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Looking_right    bool
	Looking_left     bool
	Looking_forward  bool
	Looking_backward bool

	texture  *util.Image
	hit_box  *sdl.FRect
	inventory *Inventory
}

func Create_player(texture_path string, renderer *sdl.Renderer, assets embed.FS, scale float32, init_x float32) Player {
	texture := util.Load_image(texture_path, renderer, assets, 5.0)
	hit_box := sdl.FRect{}
	return Player{Looking_forward: true, texture: &texture, hit_box: &hit_box, inventory: nil}
}
