package xavatar

import (
	"image"

	"github.com/evercyan/brick/ximg"
)

// xBase ...
type xBase struct {
	opt *option
}

func (t *xBase) Image(char string) (image.Image, error) {
	return nil, nil
}

func (t *xBase) Generate(char string) (image.Image, error) {
	return nil, nil
}

func (t *xBase) Save(char, fpath string) error {
	img, err := t.Image(char)
	if err != nil {
		return err
	}
	return ximg.Write(fpath, img)
}

// Radious 设置圆角
func (t *xBase) Radious(img image.Image, radious float64) image.Image {
	return ximg.Rounded(img, radious)
}
