package util

import (
	"embed"
	"go_outside/lib/logger"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	renderer *sdl.Renderer
	Texture  *sdl.Texture
	Surface  *sdl.Surface
	Font     *ttf.Font
}

func Load_font(font_path string, font_size int32, renderer *sdl.Renderer, assets embed.FS) Font {
	error_font := Font{renderer: nil, Texture: nil, Surface: nil, Font: nil}

	// Failed to read file
	font_data, err := assets.ReadFile(font_path)
	if err != nil {
		logger.Log("Failed to load font: "+err.Error(), logger.ERROR)
		return error_font
	}

	// Convert file to bytes
	rwops, err := sdl.RWFromMem(font_data)
	if err != nil {
		logger.Log("Failed to load image: "+sdl.GetError().Error(), logger.ERROR)
		return error_font
	}
	defer rwops.Close()

	font, err := ttf.OpenFontRW(rwops, 0, int(font_size))
	if err != nil {
		logger.Log("Failed to open font: "+err.Error(), logger.ERROR)
	}
	surface, err := font.RenderUTF8Blended("Hello, SDL2 TTF!", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		logger.Log("Failed to render text: "+err.Error(), logger.ERROR)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		logger.Log("Failed to create texture: "+err.Error(), logger.ERROR)
	}
	return Font{Font: font, renderer: renderer, Texture: texture, Surface: surface}
}

func (f *Font) Draw() {
	f.renderer.Copy(f.Texture, nil, nil)
}
