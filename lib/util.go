package lib

import (
	"embed"
	"strings"
	"fmt"

	"go_outside/lib/logger"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func Load_image(filepath string, renderer *sdl.Renderer, assets embed.FS) (*sdl.Texture, error) {
	
	// Failed to read file
	imageData, err := assets.ReadFile(filepath)
	if err != nil {
		logger.Log("Failed to load image: " + err.Error(), logger.ERROR)
		return nil, fmt.Errorf("could not load image: %v", err)
	}

	// Convert file to bytes
	rwops, err := sdl.RWFromMem(imageData)
	if err != nil {
		logger.Log("Failed to load image: " + sdl.GetError().Error(), logger.ERROR)
		return nil, fmt.Errorf("Failed to convert img to bytes: %v", err)
	}
	defer rwops.Close()

	// Load from memory
	surface_raw, err := img.LoadPNGRW(rwops)
	if err != nil {
		logger.Log("Failed to load texture from raw: " + err.Error(), logger.ERROR)
		return nil, fmt.Errorf("Failed to load texture from raw: %v", err)
	}
	defer surface_raw.Free()

	// Create the texture
	texture, err := renderer.CreateTextureFromSurface(surface_raw)
	if err != nil {
		logger.Log("Could not create texture: " + err.Error(), logger.ERROR)
		return nil, fmt.Errorf("could not create texture: %v", err)
	}

	trimmed_path := strings.Split(filepath, "/")

	logger.Log("successfully loaded texture: "+ trimmed_path[len(trimmed_path)-1], logger.SUCCESS)
	
	return texture, nil
}
