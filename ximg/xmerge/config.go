package xmerge

import "image/color"

// ...
var (
	mergeMaxRowCount   = 1000  // 最大行数
	mergeMaxColCount   = 1000  // 最大列数
	mergeMaxImageCount = 10000 // 最大图片数
	mergeDefaultOption = &Option{
		Color:   color.White,
		Padding: 20,
		Space:   10,
		Quality: 100,
	}
)
