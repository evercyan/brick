package xlogo

import (
	"fmt"
	"image"
	"image/color"
	"unicode/utf8"

	"github.com/disintegration/letteravatar"
	"github.com/evercyan/brick/xfont"
	"github.com/evercyan/brick/ximg"
	"github.com/evercyan/brick/ximg/xdraw"
	"github.com/evercyan/brick/xregex"
	"github.com/evercyan/brick/xutil"
)

// styleSingle ...
type styleSingle struct {
	*styleBase
}

// Image ...
func (t *styleSingle) Image(char string) (image.Image, error) {
	if t.opt.Width != t.opt.Height {
		return nil, fmt.Errorf("width and height must be equal")
	}
	if xutil.Len(char) != 1 {
		return nil, fmt.Errorf("char must contain only one letter")
	}
	opt := &letteravatar.Options{
		Palette:     []color.Color{t.opt.BgColor},
		LetterColor: t.opt.Color,
	}
	if xregex.HasChinese(char) {
		opt.Font = xfont.LoadAlimamaFangYuanTiVF()
	}
	firstLetter, _ := utf8.DecodeRuneInString(char)
	img, err := letteravatar.Draw(t.opt.Width, firstLetter, opt)
	if err != nil {
		return nil, err
	}
	if t.opt.Radious > 0 {
		img = xdraw.Rounded(img, t.opt.Radious)
	}
	return img, nil
}

// Save ...
func (t *styleSingle) Save(char, fpath string) error {
	img, err := t.Image(char)
	if err != nil {
		return err
	}
	return ximg.Write(fpath, img)
}

// ----------------------------------------------------------------

func newStyleSingle(opt *option) Logo {
	return &styleSingle{
		styleBase: &styleBase{
			opt: opt,
		},
	}
}
