package lib

import (
	"fmt"

	"go_outside/lib/logger"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func Load_image(filepath string, renderer *sdl.Renderer) (*sdl.Texture, error) {

	// Load the image
	image, err := img.Load("")
	if err != nil {
        return nil, fmt.Errorf("could not load image: %v", err)
    }
    defer image.Free()

	// Create the texture
	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, fmt.Errorf("could not create texture: %v", err)
	}

	logger.Log("succesfully loaded texture: %s" + filepath, logger.SUCCESS)
	return texture, nil
}