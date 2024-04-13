package xmerge

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/evercyan/brick/ximg"
)

// Merge 拼接图片
func Merge(images []image.Image, row int, col int, options ...func(*Option)) (image.Image, error) {
	if len(images) == 0 {
		return nil, fmt.Errorf("invalid images")
	}
	if row == 1 {
		return MergeToRow(images, options...)
	}
	if col == 1 {
		return MergeToCol(images, options...)
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
		resizeImg := ximg.Resize(img, width, height)
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

// MergeToRow 合并成一行, 需要满足图片的高度一致
func MergeToRow(images []image.Image, options ...func(*Option)) (image.Image, error) {
	opt := mergeDefaultOption
	opt.Space = 0
	for _, fn := range options {
		fn(opt)
	}
	// 取图片总宽, 校验图片高度
	height, sumWidth := images[0].Bounds().Dy(), 0
	for _, img := range images {
		if height != img.Bounds().Dy() {
			return nil, fmt.Errorf("image height should same when row is 1")
		}
		sumWidth += img.Bounds().Dx()
	}
	// 计算总画布宽: 边距*2 + (图片数量-1)*间距 + 图片总宽
	// 计算总画布高: 边距*2 + 图片高度
	totalWidth := opt.Padding*2 + (len(images)-1)*opt.Space + sumWidth
	totalHeight := opt.Padding*2 + height
	// 创建背景图并设置背景色
	dstImg := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	for x := 0; x < dstImg.Bounds().Dx(); x++ {
		for y := 0; y < dstImg.Bounds().Dy(); y++ {
			dstImg.Set(x, y, opt.Color)
		}
	}
	preWidth := 0
	targetH := height * opt.Quality / 100
	for index, img := range images {
		// 图片左上坐标计算: 边距 + 图片索引*间距 + 图片索引*图片宽高
		x0 := opt.Padding + index*opt.Space + preWidth
		y0 := opt.Padding
		// 将原图重置大小
		targetW := img.Bounds().Dx() * opt.Quality / 100
		resizeImg := ximg.Resize(img, targetW, targetH)
		// 画图
		draw.Draw(
			dstImg,
			image.Rect(
				x0,         // 左上 x
				y0,         // 左上 y
				x0+targetW, // 右下 x
				y0+height,  // 右下 y
			),
			resizeImg,
			image.Point{},
			draw.Over,
		)
		preWidth += targetW
	}
	return dstImg, nil
}

// MergeToCol 合并成一列, 需要满足图片的宽度一致
func MergeToCol(images []image.Image, options ...func(*Option)) (image.Image, error) {
	opt := mergeDefaultOption
	opt.Space = 0
	for _, fn := range options {
		fn(opt)
	}
	// 取图片总高, 校验图片宽度
	width, sumHeight := images[0].Bounds().Dx(), 0
	for _, img := range images {
		if width != img.Bounds().Dx() {
			return nil, fmt.Errorf("image width should same when col is 1")
		}
		sumHeight += img.Bounds().Dy()
	}
	// 计算总画布宽: 边距*2 + 图片宽度
	// 计算总画布高: 边距*2 + (图片数量-1)*间距 + 图片总高
	totalWidth := opt.Padding*2 + width
	totalHeight := opt.Padding*2 + (len(images)-1)*opt.Space + sumHeight
	// 创建背景图并设置背景色
	dstImg := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	for x := 0; x < dstImg.Bounds().Dx(); x++ {
		for y := 0; y < dstImg.Bounds().Dy(); y++ {
			dstImg.Set(x, y, opt.Color)
		}
	}
	preHeight := 0
	targetW := width * opt.Quality / 100
	for index, img := range images {
		// 图片左上坐标计算: 边距 + 图片索引*间距 + 图片索引*图片宽高
		x0 := opt.Padding
		y0 := opt.Padding + index*opt.Space + preHeight
		// 将原图重置大小
		targetH := img.Bounds().Dy() * opt.Quality / 100
		resizeImg := ximg.Resize(img, targetW, targetH)
		// 画图
		draw.Draw(
			dstImg,
			image.Rect(
				x0,         // 左上 x
				y0,         // 左上 y
				x0+width,   // 右下 x
				y0+targetH, // 右下 y
			),
			resizeImg,
			image.Point{},
			draw.Over,
		)
		preHeight += targetH
	}
	return dstImg, nil
}
