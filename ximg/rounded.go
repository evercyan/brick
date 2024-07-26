package ximg

import (
	"image"
	"math"

	"github.com/evercyan/brick/xlodash"
	"github.com/fogleman/gg"
)

// Rounded 圆角处理
// p: 圆角率, 范围 0 - 1, 默认 0.25
func Rounded(src image.Image, p float64) image.Image {
	if p < 0 || p > 1 {
		p = 0.25
	}
	w, h := float64(src.Bounds().Dx()), float64(src.Bounds().Dy())
	// 取长宽较小者的一半作为圆最小半径 minr, 取其平方根作为最大半径 maxr(即正方形的中心到顶点距离)
	// 则半径 r = maxr - (maxr-minr)*p
	minr := xlodash.Min[float64](w/2, h/2)
	maxr := math.Sqrt(minr)
	r := maxr - (maxr-minr)*p
	dc := gg.NewContext(int(w), int(h))
	dc.DrawRoundedRectangle(0, 0, w, h, r)
	dc.Clip()
	dc.DrawImage(src, 0, 0)
	return dc.Image()
}

// ----------------------------------------------------------------

// Circle 圆形
func Circle(src image.Image) image.Image {
	return Rounded(src, 1)
}

// ----------------------------------------------------------------

// Border 外边框
func Border(src image.Image, width int, colors ...string) image.Image {
	if width <= 0 {
		return src
	}
	w, h := src.Bounds().Dx(), src.Bounds().Dy()
	newW, newH := w+2*width, h+2*width
	dc := gg.NewContext(newW, newH)
	dc.DrawRectangle(0, 0, float64(newW), float64(newH))
	color := "#fff"
	if len(colors) > 0 {
		color = colors[0]
	}
	dc.SetHexColor(color)
	dc.Fill()
	dc.DrawImageAnchored(src, newW/2, newH/2, 0.5, 0.5)
	return dc.Image()
}
