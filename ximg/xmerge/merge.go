package xmerge

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/evercyan/brick/ximg/xdraw"
)

// Merge 拼接图片
func Merge(images []image.Image, row int, col int, options ...func(*Option)) (image.Image, error) {
	if len(images) == 0 {
		return nil, fmt.Errorf("invalid images")
	}
	if row > mergeMaxRowCount {
		return nil, fmt.Errorf("max row count is %d", mergeMaxRowCount)
	}
	if col > mergeMaxColCount {
		return nil, fmt.Errorf("max col count is %d", mergeMaxColCount)
	}
	if row*col > mergeMaxImageCount {
		return nil, fmt.Errorf("max image count is %d", mergeMaxImageCount)
	}
	opt := mergeDefaultOption
	for _, fn := range options {
		fn(opt)
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

		// 将原图重置大小
		resizeImg := xdraw.Resize(img, width, height)

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
