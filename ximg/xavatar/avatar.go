package xavatar

import (
	"image"

	"github.com/evercyan/brick/xconvert"
)

// Avatar ...
type Avatar interface {
	Image(char string) (image.Image, error)
	Generate(char string) (image.Image, error)
	Save(char, fpath string) error
}

// New ...
func New(options ...Option) Avatar {
	opt := getDefaultOption()
	for _, fn := range options {
		fn(opt)
	}
	switch opt.Style {
	case StyleLetter:
		return newLetter(opt)
	case StyleSquare:
		return newSquare(opt)
	default:
		panic("invalid style")
	}
}

// getDefaultOption ...
func getDefaultOption() *option {
	return &option{
		Style:    StyleLetter,
		Size:     512,
		Radious:  0,
		Color:    xconvert.Hex2Color("#000000"),
		BgColor:  xconvert.Hex2Color("#ffffff"),
		BlockNum: 9,
	}
}
