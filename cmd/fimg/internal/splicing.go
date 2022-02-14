package internal

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"regexp"
	"sort"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/ximage"
	"github.com/spf13/cobra"
)

// files 水印图片
//go:embed splicing_logo.png
var files embed.FS

var (
	ImageRegex  = `(\d+)\.(jpg|jpeg|png)$` // 图片匹配规则
	ImageOutput = "./splicing.png"         // 图片输出路径
)

var (
	SplicingRowCount  int    // 行数
	SplicingColCount  int    // 列数
	SplicingColor     string // 颜色
	SplicingPadding   int    // 边距
	SplicingSpace     int    // 间距
	SplicingQuality   int    // 质量
	SplicingWaterMark bool   // 水印
	SplicingImageDir  string // 目录
)

var (
	// SplicingCommand ...
	SplicingCommand = &cobra.Command{
		Use:   "splicing",
		Short: "拼接图片",
		Run: func(cmd *cobra.Command, args []string) {
			if SplicingImageDir == "" {
				xcolor.Fail("Error:", "图片目录不能为空")
				return
			}
			if !xfile.IsExist(SplicingImageDir) {
				xcolor.Fail("Error:", "图片目录不存在")
				return
			}
			fileList := xfile.ListFiles(SplicingImageDir, ImageRegex)
			if len(fileList) == 0 {
				xcolor.Fail("Error:", "目录下图片文件格式需要满足 xxx_数字.(png|jpg|jpeg)")
				return
			}

			// 取文件名中的数字, 进行排序, 按序拼接
			imageNums := make([]int, 0)
			imageMap := make(map[int]string)
			re := regexp.MustCompile(ImageRegex)
			for _, filePath := range fileList {
				matchs := re.FindStringSubmatch(filePath)
				if len(matchs) != 3 {
					continue
				}
				num := int(xconvert.ToUint(matchs[1]))
				if _, ok := imageMap[num]; ok {
					continue
				}
				imageMap[num] = filePath
				imageNums = append(imageNums, num)
			}
			sort.Ints(imageNums)
			images := make([]image.Image, 0)
			for _, num := range imageNums {
				img, _, err := ximage.Parse(imageMap[num])
				if err != nil {
					continue
				}
				images = append(images, img)
			}

			// 拼接图片
			options := []func(option *ximage.SplicingOption){
				ximage.WithSplicingColor(SplicingColor),     // 颜色
				ximage.WithSplicingPadding(SplicingPadding), // 边距
				ximage.WithSplicingSpace(SplicingSpace),     // 间距
				ximage.WithSplicingQuality(SplicingQuality), // 质量
			}
			// 水印
			if SplicingWaterMark {
				b, err := files.ReadFile("splicing_logo.png")
				if err == nil {
					wmImg, _, _ := image.Decode(bytes.NewReader(b))
					options = append(options, ximage.WithSplicingWaterMark(wmImg))
				}
			}
			dstImg, err := ximage.Splicing(
				images,
				SplicingRowCount,
				SplicingColCount,
				options...,
			)
			if err != nil {
				xcolor.Fail("Error:", err.Error())
				return
			}

			// 保存图片
			if err := ximage.Write(ImageOutput, dstImg); err != nil {
				xcolor.Fail("Error:", err.Error())
				return
			}

			xcolor.Success("🍺🍺🍺:", fmt.Sprintf("点击打开 %s", ImageOutput))
		},
	}
)

func init() {
	flags := SplicingCommand.PersistentFlags()
	flags.IntVarP(&SplicingRowCount, "row", "", 2, "图片行数")
	flags.IntVarP(&SplicingColCount, "col", "", 2, "图片列数")
	flags.StringVarP(&SplicingColor, "color", "", "ffffff", "背景颜色")
	flags.IntVarP(&SplicingPadding, "padding", "", 20, "图片边距")
	flags.IntVarP(&SplicingSpace, "space", "", 10, "图片间距")
	flags.IntVarP(&SplicingQuality, "quality", "", 100, "图片质量: 取值范围 1-100")
	flags.BoolVarP(&SplicingWaterMark, "watermark", "w", false, "图片水印")
	flags.StringVarP(
		&SplicingImageDir,
		"image dir",
		"d",
		"",
		"图片目录, 图片文件格式需要满足 xxx_数字.(png|jpg|jpeg)",
	)
}
