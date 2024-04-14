package xavatar

import (
	"bytes"
	"image"
	"image/color"

	"github.com/cespare/xxhash/v2"
	"github.com/evercyan/brick/ximg"
)

// xSquare ...
type xSquare struct {
	*xBase
}

// Generate ...
func (t *xSquare) Generate(char string) (image.Image, error) {
	digest := xxhash.Sum64String(char)
	// 根据 digest 计算颜色
	img := image.NewPaletted(image.Rect(0, 0, t.opt.Size, t.opt.Size), color.Palette{
		color.NRGBA{R: byte(digest), G: byte(digest >> 8), B: byte(digest >> 16), A: 0xff},
		color.NRGBA{R: 0xff ^ byte(digest), G: 0xff ^ byte(digest>>8), B: 0xff ^ byte(digest>>16), A: 0xff},
	})
	// 方块数量
	blockNum := t.opt.BlockNum
	// 单个方块尺寸
	blockSize := t.opt.Size / (blockNum + 1)
	// 计算边距 = 方块尺寸一半 + 多作空白一半
	padding := blockSize/2 + (t.opt.Size%blockSize)/2
	filled := blockNum == 1
	pixels := bytes.Repeat([]byte{1}, blockSize)
	for i, ri, ci := 0, 0, 0; i < blockNum*(blockNum+1)/2; i++ {
		if filled || digest>>uint(i%64)&1 == 1 {
			for i := 0; i < blockSize; i++ {
				x := padding + ri*blockSize
				y := padding + ci*blockSize + i
				copy(img.Pix[img.PixOffset(x, y):], pixels)
				x = padding + (blockNum-1-ri)*blockSize
				copy(img.Pix[img.PixOffset(x, y):], pixels)
			}
		}
		ci++
		if ci == blockNum {
			ci = 0
			ri++
		}
	}
	var targetImg image.Image = img
	if t.opt.Radious > 0 {
		targetImg = ximg.Rounded(img, t.opt.Radious)
	}
	return targetImg, nil
}

// Save ...
func (t *xSquare) Save(char, fpath string) error {
	img, err := t.Generate(char)
	if err != nil {
		return err
	}
	return ximg.Write(fpath, img)
}

// ----------------------------------------------------------------

func newSquare(opt *option) Avatar {
	return &xSquare{
		xBase: &xBase{
			opt: opt,
		},
	}
}
