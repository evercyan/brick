package ximage

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/evercyan/brick/xlodash"
)

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.p.X, c.p.Y)
}

func (c *circle) At(x, y int) color.Color {
	var xx, yy, rr float64
	var inArea bool
	if x <= c.r && y <= c.r {
		xx, yy, rr = float64(c.r-x)+0.5, float64(y-c.r)+0.5, float64(c.r)
		inArea = true
	}
	if x >= (c.p.X-c.r) && y <= c.r {
		xx, yy, rr = float64(x-(c.p.X-c.r))+0.5, float64(y-c.r)+0.5, float64(c.r)
		inArea = true
	}
	if x <= c.r && y >= (c.p.Y-c.r) {
		xx, yy, rr = float64(c.r-x)+0.5, float64(y-(c.p.Y-c.r))+0.5, float64(c.r)
		inArea = true
	}
	if x >= (c.p.X-c.r) && y >= (c.p.Y-c.r) {
		xx, yy, rr = float64(x-(c.p.X-c.r))+0.5, float64(y-(c.p.Y-c.r))+0.5, float64(c.r)
		inArea = true
	}
	if inArea && xx*xx+yy*yy >= rr*rr {
		return color.Alpha{}
	}
	return color.Alpha{A: 255}
}

// Radius ...
func Radius(img image.Image, r int) image.Image {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	// 圆角半径不超过 img 宽或高一半
	r = xlodash.Min[int](w/2, h/2, r)
	c := &circle{
		p: image.Point{
			X: w,
			Y: h,
		},
		r: r,
	}
	radiusImg := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.DrawMask(
		radiusImg,
		radiusImg.Bounds(),
		img,
		image.Point{},
		c,
		image.Point{},
		draw.Over,
	)
	return radiusImg
}
