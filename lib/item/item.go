package item

import (
	"embed"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Items struct {
	renderer  *sdl.Renderer
	assets    embed.FS
	font_path string

	registered []*Item
}

type Item struct {
	font    util.Font
	Texture *util.Image
	Name    string
	Rarity  int32
}

// Initialzies the items registry
func Init_items(renderer *sdl.Renderer, assets embed.FS, font_path string) *Items {
	return &Items{renderer: renderer, assets: assets, font_path: font_path}
}

// This is used to register an item
// you can also use it as a reference
// to this item by declaring it as a variable
func (i *Items) New(name string, rarity int32) *Item {
	texture := util.Load_image("assets/textures/items/"+name, i.renderer, i.assets, 2.0)
	texture.X, texture.Y = 300, 300
	font := util.Load_font(i.font_path, 12, name, &sdl.Color{R: 255, G: 255, B: 255}, i.renderer, i.assets)

	item := &Item{Name: name, Rarity: rarity, Texture: &texture, font: font}
	i.registered = append(i.registered, item)
	return item
}

// renders all items registered in the
// the Items struct
func (i *Items) Draw() {
	for _, item := range i.registered {
		println(item.Name)
	}

}

func (i *Item) Draw_single(x, y *float32) {
	i.Texture.Draw_image(x, y)
	x_pos := i.Texture.X - float32(i.font.Surface.W) / 3.0
	i.font.Draw(x_pos, i.Texture.Y-20)
}
