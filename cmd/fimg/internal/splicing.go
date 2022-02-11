package internal

import (
	"image"
	"image/png"
	"os"
	"regexp"
	"sort"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/ximage"
	"github.com/spf13/cobra"
)

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
				xcolor.Fail("图片路径为空:", "e.g. fimg join /tmp/xxx")
				return
			}
			if !xfile.IsExist(SplicingImageDir) {
				xcolor.Fail("图片路径无效:", SplicingImageDir)
				return
			}
			fileList := xfile.ListFiles(SplicingImageDir, ImageRegex)
			if len(fileList) == 0 {
				xcolor.Fail("未找到有效的图片文件:", SplicingImageDir)
				return
			}

			// 截取文件名中的数字, 进行排序, 方便按序拼接
			imageMap := make(map[int]string)
			nums := make([]int, 0)
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
				nums = append(nums, num)
			}
			sort.Ints(nums)

			images := make([]image.Image, 0)
			for _, num := range nums {
				img, _, err := ximage.Parse(imageMap[num])
				if err != nil {
					continue
				}
				images = append(images, img)
			}

			// 拼接图片
			dstImg, err := ximage.Splicing(
				images,
				SplicingRowCount,
				SplicingColCount,
				ximage.WithSplicingColor(SplicingColor),
				ximage.WithSplicingPadding(SplicingPadding),
				ximage.WithSplicingSpace(SplicingSpace),
				ximage.WithSplicingQuality(SplicingQuality),
				ximage.WithSplicingWaterMark(SplicingWaterMark),
			)
			if err != nil {
				xcolor.Fail("拼接图片失败:", err.Error())
				return
			}

			// 保存图片
			file, _ := os.Create(ImageOutput)
			defer file.Close()
			if err := png.Encode(file, dstImg); err != nil {
				xcolor.Fail("拼接图片失败:", err.Error())
				return
			}
			xcolor.Success("拼接图片成功:", ImageOutput)
		},
	}
)

func init() {
	flags := SplicingCommand.PersistentFlags()
	flags.IntVarP(&SplicingRowCount, "row", "", 2, "图片行数")
	flags.IntVarP(&SplicingColCount, "col", "", 2, "图片列数")
	flags.StringVarP(&SplicingColor, "color", "", "#ffffff", "背景颜色")
	flags.IntVarP(&SplicingPadding, "padding", "", 20, "图片边距")
	flags.IntVarP(&SplicingSpace, "space", "", 10, "图片间距")
	flags.IntVarP(&SplicingQuality, "quality", "", 100, "图片质量: 取值范围 1-100")
	flags.BoolVarP(&SplicingWaterMark, "watermark", "w", true, "图片水印")
	flags.StringVarP(&SplicingImageDir, "image dir", "d", "", "图片目录, 图片文件格式需要满足 xxx_数字.(png|jpg|jpeg)")
}
