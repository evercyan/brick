package xlogo

import (
	"image"
)

// Logo ...
type Logo interface {
	Image(char string) (image.Image, error)
	Save(char, fpath string) error
}

// New ...
func New(options ...Option) Logo {
	opt := defaultOption
	for _, fn := range options {
		fn(opt)
	}
	switch opt.Style {
	case StyleSingle:
		return newStyleSingle(opt)
	default:
		panic("invalid style")
	}
}
