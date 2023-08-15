package item

import (
	"embed"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Items struct {
	Renderer  *sdl.Renderer
	Assets    embed.FS
	font_path string

	registered []*Item
}

type Item struct {
	font     util.Font
	Texture  *util.Image
	rect 	 *sdl.Rect
	Name     string
	settings *Settings
}

type Settings struct {
	Rarity     int32
	Durability int32
	Tooltip    string
}

// Initialzies the items registry
func InitItems(renderer *sdl.Renderer, assets embed.FS, font_path string) *Items {
	return &Items{Renderer: renderer, Assets: assets, font_path: font_path}
}

// This is used to register an item
// you can also use it as a reference
// to this item by declaring it as a variable
func (i *Items) New(name string, settings Settings) *Item {
	texture := util.LoadImage("assets/textures/items/"+name, i.Renderer, i.Assets, 2.0, true)
	texture.X, texture.Y = 300, 300
	font := util.LoadFont(i.font_path, 12, name, &sdl.Color{R: 255, G: 255, B: 255}, i.Renderer, i.Assets)
	rect := util.Convert_frect_to_rect(texture.Image_rect)

	item := &Item{Name: name, Texture: texture, font: font, rect: &rect}
	i.registered = append(i.registered, item)
	return item
}

func New() *Settings {
	return &Settings{}
}

func (is *Settings) SetRarity(rarity int32) {
	is.Rarity = rarity
}

// renders all items registered in the
// the Items struct
func Draw(i *Items) {
	for _, item := range i.registered {
		rect := sdl.Rect{X: int32(item.Texture.X), Y: int32(item.Texture.Y), W: int32(item.Texture.Image_rect.W), H: int32(item.Texture.Image_rect.H)}
		if util.MouseCollide(&rect) {
			println("collide with mouse" + item.Name)
		}
	}
}

// Draw a single item
func (i *Item) Draw_single(x, y *float32) {
	i.Texture.Draw_image(x, y)
	i.Texture.X, i.Texture.Y = *x, *y
	x_pos := i.Texture.X - float32(i.font.Surface.W)/3.0
	i.font.Draw(x_pos, i.Texture.Y-20)
}
