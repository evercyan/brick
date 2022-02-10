package internal

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xfile"
	"github.com/spf13/cobra"
)

var (
	ImageRegex = `(\d+)\.(jpg|jpeg|png)$`
)

var (
	// JoinCommand ...
	JoinCommand = &cobra.Command{
		Use:   "join",
		Short: "拼接图片",
		Run: func(cmd *cobra.Command, args []string) {

			// TODO m*n 上限 10000
			// TODO 图片数量上限

			if JoinImageDir == "" {
				xcolor.Fail("图片路径不能为空:", "e.g. fimg join /tmp/xxx")
				return
			}
			if !xfile.IsExist(JoinImageDir) {
				xcolor.Fail("图片路径无效:", JoinImageDir)
				return
			}
			fileList := xfile.ListFiles(JoinImageDir, ImageRegex)
			if len(fileList) == 0 {
				xcolor.Fail("未找到有效的图片文件:", JoinImageDir)
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
			images := make([]string, 0)
			for _, num := range nums {
				images = append(images, imageMap[num])
			}

			fmt.Println(images)
		},
	}
)

var (
	JoinRowCount int
	JoinColCount int
	JoinPxSpace  int
	JoinImageDir string
)

func init() {
	JoinCommand.PersistentFlags().IntVarP(&JoinRowCount, "image row", "r", 2, "行数")
	JoinCommand.PersistentFlags().IntVarP(&JoinColCount, "image col", "c", 2, "列数")
	JoinCommand.PersistentFlags().IntVarP(&JoinPxSpace, "image space", "s", 10, "间隔")
	JoinCommand.PersistentFlags().StringVarP(&JoinImageDir, "image dir", "d", "", "图片路径")
}
