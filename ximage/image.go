package ximage

import (
	"image"
	"image/color"

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

// Circle ...
func Circle(src image.Image) image.Image {
	d := src.Bounds().Dx()
	if src.Bounds().Dy() < d {
		d = src.Bounds().Dy()
	}
	dst := imaging.CropCenter(src, d, d)
	r := d / 2
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			if (x-r)*(x-r)+(y-r)*(y-r) > r*r {
				dst.SetNRGBA(x, y, color.NRGBA{})
			}
		}
	}
	return dst
}
