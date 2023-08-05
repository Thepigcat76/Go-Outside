package item

import (
	"embed"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Items struct {
	renderer *sdl.Renderer
	assets   embed.FS

	registered []*Item
}

type Item struct {
	Texture *util.Image
	Name    string
	Rarity  int32
}

// Initialzies the items registry
func Init_items(renderer *sdl.Renderer, assets embed.FS) *Items {
	return &Items{renderer: renderer, assets: assets}
}

// This is used to register an item
// you can also use it as a reference
// to this item by declaring it as a variable
func (i *Items) New(name string, rarity int32) *Item {
	texture := util.Load_image("assets/textures/items/"+name, i.renderer, i.assets, 2.0)

	item := &Item{Name: name, Rarity: rarity, Texture: &texture}
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

func (i *Item) Draw_single() {
	i.Texture.Draw_image()
}
