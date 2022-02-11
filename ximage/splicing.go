package ximage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/evercyan/brick/xconvert"
)

var (
	maxRowCount           = 1000  // 最大行数
	maxColCount           = 1000  // 最大列数
	maxImageCount         = 10000 // 最大图片数
	defaultSplicingOption = &SplicingOption{
		Color:     color.White,
		Padding:   20,
		Space:     10,
		Quality:   100,
		WaterMark: true,
	}
)

// SplicingOption 拼接选项
type SplicingOption struct {
	Color     color.Color // 背景颜色
	Padding   int         // 上右下左边距
	Space     int         // 图片间距
	Quality   int         // 图片质量
	WaterMark bool        // 水印开关
}

// WithSplicingColor 背景颜色
func WithSplicingColor(v string) func(option *SplicingOption) {
	return func(option *SplicingOption) {
		option.Color = xconvert.Hex2Color(v)
	}
}

// WithSplicingPadding 边距, 1-100
func WithSplicingPadding(v int) func(option *SplicingOption) {
	return func(option *SplicingOption) {
		if v >= 1 && v <= 100 {
			option.Padding = v
		}
	}
}

// WithSplicingSpace 图片间距, 1-100
func WithSplicingSpace(v int) func(option *SplicingOption) {
	return func(option *SplicingOption) {
		if v >= 1 && v <= 100 {
			option.Space = v
		}
	}
}

// WithSplicingQuality 图片质量, 默认 100, 取值范围 1-100
func WithSplicingQuality(v int) func(option *SplicingOption) {
	return func(option *SplicingOption) {
		if v >= 1 && v <= 100 {
			option.Quality = v
		}
	}
}

// WithSplicingWaterMark 水印
func WithSplicingWaterMark(v bool) func(option *SplicingOption) {
	return func(option *SplicingOption) {
		option.WaterMark = v
	}
}

// ----------------------------------------------------------------

// Splicing ...
func Splicing(
	images []image.Image,
	row int,
	col int,
	options ...func(*SplicingOption),
) (image.Image, error) {
	if len(images) == 0 {
		return nil, fmt.Errorf("invalid images")
	}
	if row > maxRowCount {
		return nil, fmt.Errorf("max row count is %d", maxRowCount)
	}
	if col > maxColCount {
		return nil, fmt.Errorf("max col count is %d", maxColCount)
	}
	if row*col > maxImageCount {
		return nil, fmt.Errorf("max image count is %d", maxImageCount)
	}
	opt := defaultSplicingOption
	for _, option := range options {
		option(opt)
	}

	// 取图片平均宽度
	sumWidth, sumHeight := 0, 0
	for _, img := range images {
		sumWidth += img.Bounds().Dx()
		sumHeight += img.Bounds().Dy()
	}
	width, height := sumWidth/len(images), sumHeight/len(images)

	// 根据图片质量缩小宽高
	width, height = width*opt.Quality/100, height*opt.Quality/100

	// 计算总画布宽高: 边距*2 + (图片数量-1)*间距 + 图片数量*图片宽高
	totalWidth := opt.Padding*2 + (col-1)*opt.Space + width*col
	totalHeight := opt.Padding*2 + (row-1)*opt.Space + height*row
	// 创建背景图并设置背景色
	dstImg := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	// 设置背景色
	for x := 0; x < dstImg.Bounds().Dx(); x++ {
		for y := 0; y < dstImg.Bounds().Dy(); y++ {
			dstImg.Set(x, y, opt.Color)
		}
	}

	// 总需要图片数量
	total := row * col
	for index, img := range images {
		if index >= total {
			break
		}
		rowIndex := index % col
		colIndex := index / col

		// 图片左上坐标计算: 边距 + 图片索引*间距 + 图片索引*图片宽高
		x0 := opt.Padding + rowIndex*opt.Space + rowIndex*width
		y0 := opt.Padding + colIndex*opt.Space + colIndex*height

		// TODO debug
		fmt.Printf(
			"index: %d, rowIndex: %d, colIndex: %d, point: (%d, %d)\n",
			index, rowIndex, colIndex, x0, y0,
		)

		// 将原图重置大小
		resizeImg := Resize(img, width, height)

		// 画图
		draw.Draw(
			dstImg,
			image.Rect(
				x0,        // 左上 x
				y0,        // 左上 y
				x0+width,  // 右下 x
				y0+height, // 右下 y
			),
			resizeImg,
			image.Point{},
			draw.Over,
		)
	}
	return dstImg, nil
}
