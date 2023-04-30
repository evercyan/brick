package xmerge

import (
	"image/color"

	"github.com/evercyan/brick/xconvert"
)

// Option 拼接选项
type Option struct {
	Color   color.Color // 背景颜色
	Padding int         // 上右下左边距
	Space   int         // 图片间距
	Quality int         // 图片质量
}

// WithColor 背景颜色
func WithColor(v string) func(option *Option) {
	return func(option *Option) {
		option.Color = xconvert.Hex2Color(v)
	}
}

// WithPadding 边距, 1-100
func WithPadding(v int) func(option *Option) {
	return func(option *Option) {
		if v >= 1 && v <= 100 {
			option.Padding = v
		}
	}
}

// WithSpace 图片间距, 1-100
func WithSpace(v int) func(option *Option) {
	return func(option *Option) {
		if v >= 0 && v <= 100 {
			option.Space = v
		}
	}
}

// WithQuality 图片质量, 默认 100, 取值范围 1-100
func WithQuality(v int) func(option *Option) {
	return func(option *Option) {
		if v >= 1 && v <= 100 {
			option.Quality = v
		}
	}
}
