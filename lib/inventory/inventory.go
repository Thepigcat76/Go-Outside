package inventory

import (
	"embed"
	"go_outside/lib/item"
	"go_outside/lib/util"

	"github.com/veandco/go-sdl2/sdl"
)

type Slot struct {
	slot_content  *item.Item
	content_count int32
}

func register_slot() Slot {
	return Slot{content_count: 0}
}

type Inventory struct {
	texture    *util.Image
	Slot_count int32
	Slots      []Slot
}

func (i *Inventory) Set_item(item item.Item, slot int32) {
	i.Slots[slot].slot_content = &item
}

func Init_inventory(renderer *sdl.Renderer, assets embed.FS) *Inventory {
	slot1, slot2, slot3 := register_slot(), register_slot(), register_slot()
	slots := []Slot{slot1, slot2, slot3}
	texture := util.Load_image("assets/textures/slot", renderer, assets, 5.0)
	return &Inventory{Slot_count: 3, texture: &texture, Slots: slots}
}

func (i *Inventory) Draw(X, Y float32) {
    for y := 0; y < int(i.Slot_count); y++ {
        i.texture.X, i.texture.Y = X, Y
        i.texture.Y += float32(y) * 80.0
        i.texture.Draw_image(nil, nil)

        slot := &i.Slots[y]
        if slot.slot_content != nil {
            slot.slot_content.Draw_single(&i.texture.X, &i.texture.Y)
        }
    }
}