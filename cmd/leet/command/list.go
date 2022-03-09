package command

import (
	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/cmd/leet/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/spf13/cobra"
)

var (
	// QuestionCommand ...
	ListCommand = &cobra.Command{
		Use:   "list",
		Short: "查询问题列表, e.g. leet list two-sum",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := internal.NewApp().GetList(args...)
			if err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(res)
		},
	}
)
