package ximg

import (
	"image"

	"github.com/disintegration/imaging"
)

// Resize ...
func Resize(src image.Image, w, h int) image.Image {
	if w == 0 && h == 0 {
		return src
	} else if src.Bounds().Dy() == w && src.Bounds().Dy() == h {
		return src
	} else if w == 0 {
		w = src.Bounds().Dx() * h / src.Bounds().Dy()
	} else if h == 0 {
		h = src.Bounds().Dy() * w / src.Bounds().Dx()
	}
	return imaging.Resize(src, w, h, imaging.NearestNeighbor)
}
