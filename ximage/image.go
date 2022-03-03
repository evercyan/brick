package ximage

import (
	"image"

	"github.com/disintegration/imaging"
	"github.com/noelyahan/mergi"
)

// Resize ...
func Resize(img image.Image, width, height int) image.Image {
	return imaging.Resize(img, width, height, imaging.NearestNeighbor)
}

// WaterMarkImage ...
func WaterMarkImage(img, wmImg image.Image, p image.Point) image.Image {
	res, err := mergi.Watermark(wmImg, img, p)
	if err != nil {
		return img
	}
	return res
}
