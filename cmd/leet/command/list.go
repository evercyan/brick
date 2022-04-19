package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/cmd/leet/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/evercyan/brick/xconvert"
	"github.com/spf13/cobra"
)

var (
	// ListCommand ...
	ListCommand = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "查询答题列表, e.g. leet list two-sum",
		Run: func(cmd *cobra.Command, args []string) {
			defer func(begin time.Time) {
				xcolor.Success(
					config.SymbolTime,
					fmt.Sprintf("耗时: %s", time.Now().Sub(begin).String()),
				)
			}(time.Now())

			list, err := internal.NewApp().GetQuestionList()
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			matchs := make([][]interface{}, 0)
			text, num := "", int64(0)
			if len(args) > 0 {
				text, num = args[0], int64(xconvert.ToUint(args[0]))
			}
			index := 1
			for _, v := range list {
				if strings.Contains(v.Fid, text) ||
					strings.Contains(v.Title, text) ||
					strings.Contains(v.Slug, text) ||
					v.Qid == num {
					matchs = append(matchs, []interface{}{
						index,
						v.Fid,
						v.Title,
						v.Slug,
						v.Level.String(),
					})
					index++
				}
			}
			xcolor.Success(xtable.New(matchs).Style(xtable.Dashed).Border(true).Header([]string{
				"#", "ID", "标题", "标识", "难度",
			}).Text())
		},
	}
)
