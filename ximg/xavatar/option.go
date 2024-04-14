package xavatar

import (
	"image/color"
)

// option ...
type option struct {
	Style    Style       // 绘制样式
	Size     int         // 尺寸
	Radious  float64     // 圆角, 范围 0-1
	Color    color.Color // 字体色
	BgColor  color.Color // 背景色
	BlockNum int         // 样式为 square 时的方块数量
}

// Option ...
type Option func(o *option)

// ----------------------------------------------------------------

// Style ...
type Style int

const (
	StyleLetter Style = iota
	StyleSquare
	StylePornhub
)

func (t Style) String() string {
	switch t {
	case StyleLetter:
		return "letter"
	case StyleSquare:
		return "square"
	case StylePornhub:
		return "pornhub"
	default:
		return ""
	}
}

// ----------------------------------------------------------------

// WithStyle 设置绘制样式
func WithStyle(style Style) Option {
	return func(o *option) {
		o.Style = style
	}
}

// WithSize 设置尺寸
func WithSize(size int) Option {
	return func(o *option) {
		if size < 100 {
			size = 100
		} else if size > 4096 {
			size = 4096
		}
		o.Size = size
	}
}

// WithRadious 设置圆角
func WithRadious(radious float64) Option {
	return func(o *option) {
		if radious >= 0 && radious <= 1.0 {
			o.Radious = radious
		}
	}
}

// WithBlockNum 设置方块数量
func WithBlockNum(blockNum int) Option {
	return func(o *option) {
		if blockNum > 0 && blockNum <= 20 {
			o.BlockNum = blockNum
		}
	}
}
