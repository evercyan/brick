package command

import (
	"fmt"
	"time"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/cmd/leet/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/spf13/cobra"
)

var (
	// TagCommand ...
	TagCommand = &cobra.Command{
		Use:     "tag",
		Aliases: []string{"t"},
		Short:   "生成答题标签文件",
		Run: func(cmd *cobra.Command, args []string) {
			defer func(begin time.Time) {
				xcolor.Success(
					config.SymbolTime,
					fmt.Sprintf("耗时: %s", time.Now().Sub(begin).String()),
				)
			}(time.Now())
			app := internal.NewApp()
			// 标签列表
			tagList, err := app.GetTagList()
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			// 标签问题 map
			tagMap, err := app.GetTagQuestionMap(app)
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			// 生成标签文件
			if err := app.GenerateTag(tagList, tagMap, app.Path); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(config.SymbolSuccess, "生成答题标签文件成功")
		},
	}
)
