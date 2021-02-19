package utils

import (
	"image"
	"image/png"
	"os"
)

// Img should be passed as a param to Decode method
type Img struct {
	Width      int
	Height     int
	ComponentX int
	ComponentY int
	Pixels     []uint8
}

// GetImg method takes component and img path and returns Img struct as a result
// works with png images only at the moment
func GetImg(componentX, componentY int, path string) (resImg *Img, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return
	}

	var pixels []uint8

	var width, height int

	switch trueim := img.(type) {
	case *image.RGBA:
		pixels = trueim.Pix
		width = trueim.Bounds().Max.X
		height = trueim.Bounds().Max.Y
	case *image.NRGBA:
		pixels = trueim.Pix
		width = trueim.Bounds().Max.X
		height = trueim.Bounds().Max.Y
	}

	return &Img{
		Pixels:     pixels,
		Width:      width,
		Height:     height,
		ComponentX: componentX,
		ComponentY: componentY,
	}, nil
}
