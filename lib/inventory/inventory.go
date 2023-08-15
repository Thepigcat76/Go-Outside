package inventory

import (
	"embed"
	"go_outside/lib/item"
	"go_outside/lib/logger"
	"go_outside/lib/util"
	"strconv"

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
	err_msg := "exceeded slot index, max slot is: " + strconv.Itoa(len(i.Slots)-1)
	if slot >= int32(len(i.Slots)) {
		slot = int32(len(i.Slots)) - 1
		logger.Log(err_msg, logger.WARNING)
	}
	i.Slots[slot].slot_content = &item
}

func Init_inventory(renderer *sdl.Renderer, assets embed.FS) *Inventory {
	slot1, slot2, slot3 := register_slot(), register_slot(), register_slot()
	slots := []Slot{slot1, slot2, slot3}
	texture := util.LoadImage("assets/textures/slot", renderer, assets, 5.0, true)
	return &Inventory{Slot_count: 3, texture: texture, Slots: slots}
}

func (i *Inventory) Draw(items *item.Items, X, Y float32) {
	for y := 0; y < int(i.Slot_count); y++ {
		i.texture.X, i.texture.Y = X, Y
		i.texture.Y += float32(y) * 80.0
		i.texture.Draw_image(nil, nil)

		slot := &i.Slots[y]
		if slot.slot_content != nil {
			display_item := util.LoadImage("assets/textures/items/"+slot.slot_content.Name, items.Renderer, items.Assets, 4, false)
			display_item.Draw_image(&i.texture.X, &i.texture.Y)
		}
	}
}
