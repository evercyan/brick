package xavatar

import (
	"fmt"
	"image"
	"image/color"
	"unicode/utf8"

	"github.com/disintegration/letteravatar"
	"github.com/evercyan/brick/xfont"
	"github.com/evercyan/brick/ximg"
	"github.com/evercyan/brick/xregex"
	"github.com/evercyan/brick/xutil"
)

// xLetter ...
type xLetter struct {
	*xBase
}

// Image ...
func (t *xLetter) Image(char string) (image.Image, error) {
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
	img, err := letteravatar.Draw(t.opt.Size, firstLetter, opt)
	if err != nil {
		return nil, err
	}
	if t.opt.Radious > 0 {
		img = ximg.Rounded(img, t.opt.Radious)
	}
	return img, nil
}

// Save ...
func (t *xLetter) Save(char, fpath string) error {
	return t.xBase.Save(char, fpath)
}

// ----------------------------------------------------------------

func newLetter(opt *option) Avatar {
	return &xLetter{
		xBase: &xBase{
			opt: opt,
		},
	}
}
