package xdraw

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/noelyahan/mergi"
	"golang.org/x/image/font/gofont/goregular"
)

// WatermarkImage ...
func WatermarkImage(src, w image.Image, p image.Point) image.Image {
	dst, err := mergi.Watermark(w, src, p)
	if err != nil {
		return src
	}
	return dst
}

// WatermarkText ...
func WatermarkText(src image.Image, text string, color string, size, x, y float64) image.Image {
	w, h := src.Bounds().Dx(), src.Bounds().Dy()
	dc := gg.NewContext(w, h)
	dc.DrawImage(src, 0, 0)
	dc.SetHexColor(color)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return src
	}
	face := truetype.NewFace(font, &truetype.Options{Size: size})
	dc.SetFontFace(face)
	dc.DrawString(text, x, y)
	return dc.Image()
}
