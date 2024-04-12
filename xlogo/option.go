package xlogo

import (
	"image/color"
)

// option ...
type option struct {
	Style   Style       // 绘制样式
	Width   int         // 宽
	Height  int         // 高
	Radious float64     // 圆角, 范围 0-1
	Color   color.Color // 字体色
	BgColor color.Color // 背景色
}

// Option ...
type Option func(o *option)

// ----------------------------------------------------------------

// Style ...
type Style int

const (
	StyleSingle Style = iota
	StyleYoutube
	StylePornhub
)

func (t Style) String() string {
	switch t {
	case StyleSingle:
		return "single"
	case StyleYoutube:
		return "youtube"
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

// WithSize 设置宽高
func WithSize(size int) Option {
	return func(o *option) {
		o.Width = size
		o.Height = size
	}
}

// WithWidth 设置宽度
func WithWidth(width int) Option {
	return func(o *option) {
		o.Width = width
	}
}

// WithHeight 设置高度
func WithHeight(height int) Option {
	return func(o *option) {
		o.Height = height
	}
}

// WithHeight 设置圆角
func WithRadious(radious float64) Option {
	return func(o *option) {
		if radious >= 0 && radious <= 1.0 {
			o.Radious = radious
		}
	}
}
