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
	// RecordCommand ...
	RecordCommand = &cobra.Command{
		Use:     "record",
		Aliases: []string{"r"},
		Short:   "生成答题纪录文件",
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
			if err := app.GenerateRecord(list, app.Path); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(config.SymbolSuccess, "生成答题纪录文件成功")
		},
	}
)
