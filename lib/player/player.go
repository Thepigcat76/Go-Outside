package player

import (
	"embed"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Looking_forward  bool
	Looking_backward bool
	Looking_right    bool
	Looking_left     bool

	Texture   *util.Image
	inventory *Inventory
	renderer  *sdl.Renderer

	X, Y float32
}

// textures: forward = 0, backward = 1, right = 2, left = 3
func Create(textures [4]string, renderer *sdl.Renderer, assets embed.FS, scale, x, y float32) Player {
	texture := util.Load_image(textures[0], renderer, assets, scale)
	return Player{Looking_forward: true, Texture: &texture, inventory: nil, renderer: renderer, X: x, Y: y}
}

func (p *Player) Draw() {
	p.Texture.X, p.Texture.Y = p.X, p.Y
	p.Texture.Draw_image()
}
