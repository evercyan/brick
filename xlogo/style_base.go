package xlogo

import (
	"image"

	"github.com/evercyan/brick/ximg/xdraw"
)

// styleBase ...
type styleBase struct {
	opt *option
}

func (t *styleBase) Image(char string) (image.Image, error) {
	return nil, nil
}

func (t *styleBase) Save(char, fpath string) error {
	return nil
}

// Radious 设置圆角
func (t *styleBase) Radious(img image.Image, radious float64) image.Image {
	return xdraw.Rounded(img, radious)
}
