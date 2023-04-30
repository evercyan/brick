package xdraw

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize ...
func Resize(src image.Image, w, h int) image.Image {
	return imaging.Resize(src, w, h, imaging.NearestNeighbor)
}
