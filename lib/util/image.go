/*
This file handles everything related to  images like loading, embeding and drawing them to the screen

It also allows you to create animations
*/

package util

import (
	"embed"
	"strings"

	"go_outside/lib/logger"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Image struct {
	Image_rect *sdl.FRect
	Texture    *sdl.Texture
	renderer   *sdl.Renderer
	X, Y       float32
}

func LoadImage(filepath string, renderer *sdl.Renderer, assets embed.FS, scale float32, enable_logger bool) *Image {

	error_image := &Image{Image_rect: nil, Texture: nil, renderer: renderer}

	// Failed to read file
	image_data, err := assets.ReadFile(filepath + ".png")
	if err != nil {
		logger.Log("Failed to load image: "+err.Error(), logger.ERROR)
		return error_image
	}

	// Convert file to bytes
	rwops, err := sdl.RWFromMem(image_data)
	if err != nil {
		logger.Log("Failed to load image: "+sdl.GetError().Error(), logger.ERROR)
		return error_image
	}
	defer rwops.Close()

	// Load from memory
	surface_raw, err := img.LoadPNGRW(rwops)
	if err != nil {
		logger.Log("Failed to load texture from raw: "+err.Error(), logger.ERROR)
		return error_image
	}
	defer surface_raw.Free()

	image_rect := sdl.FRect{X: 0.0, Y: 0.0, W: float32(surface_raw.W) * scale, H: float32(surface_raw.H) * scale}

	// Create the texture
	texture, err := renderer.CreateTextureFromSurface(surface_raw)
	if err != nil {
		logger.Log("Could not create texture: "+err.Error(), logger.ERROR)
		return error_image
	}

	trimmed_path := strings.Split(filepath, "/")

	if enable_logger {
		logger.Log("successfully loaded texture: "+trimmed_path[len(trimmed_path)-1], logger.SUCCESS)
	}

	return &Image{Texture: texture, Image_rect: &image_rect, renderer: renderer, X: 0, Y: 0}
}

func (i Image) Draw_image(x, y *float32) {
	if x == nil {
		i.Image_rect.X = i.X
	} else {
		i.Image_rect.X = *x
	}
	if y == nil {
		i.Image_rect.Y = i.Y
	} else {
		i.Image_rect.Y = *y
	}
	texture_rect := sdl.Rect{X: int32(i.Image_rect.X), Y: int32(i.Image_rect.Y), W: int32(i.Image_rect.W), H: int32(i.Image_rect.H)}
	i.renderer.Copy(i.Texture, nil, &texture_rect)
}
