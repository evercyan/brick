package ximage

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize ...
func Resize(img image.Image, width, height int) image.Image {
	return imaging.Resize(img, width, height, imaging.NearestNeighbor)
}
