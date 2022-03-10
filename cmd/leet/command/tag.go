package command

import (
	"fmt"
	"sync"
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
			list, err := app.GetQuestionList()
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			var wg sync.WaitGroup
			for k, v := range list {
				wg.Add(1)
				go func(k int, slug string) {
					defer wg.Done()
					detail, err := app.GetQuestionDetail(slug)
					if err != nil {
						return
					}
					list[k].Detail = detail
				}(k, v.Slug)
			}
			wg.Wait()
			tagList, err := app.GetTagList()
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			if err := app.GenerateTag(list, tagList, app.Path); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(config.SymbolSuccess, "生成答题标签文件成功")
		},
	}
)
