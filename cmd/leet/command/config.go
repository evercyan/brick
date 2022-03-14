package command

import (
	"github.com/evercyan/brick/cmd/leet/config"
	"github.com/evercyan/brick/cmd/leet/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/spf13/cobra"
)

var (
	// ConfigCommand ...
	ConfigCommand = &cobra.Command{
		Use:     "config",
		Aliases: []string{"c"},
		Short:   "设置答题配置",
		Run: func(cmd *cobra.Command, args []string) {
			xcolor.Success(internal.NewApp().Render())
		},
	}
	// ConfigPathCommand ...
	ConfigPathCommand = &cobra.Command{
		Use:   "path",
		Short: "答题配置: 设置答题文件生成目录",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail(config.SymbolError, "请输入答题文件生成目录, e.g. leet config path ~/leet")
				return
			}
			app := internal.NewApp()
			if err := app.SetPath(args[0]); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(app.Render())
		},
	}
	// ConfigLangCommand ...
	ConfigLangCommand = &cobra.Command{
		Use:   "lang",
		Short: "答题配置: 设置答题文件默认语言",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail(config.SymbolError, "答题文件默认语言, e.g. leet config lang golang")
				return
			}
			app := internal.NewApp()
			if err := app.SetLang(args[0]); err != nil {
				xcolor.Fail(config.SymbolError, err.Error())
				return
			}
			xcolor.Success(app.Render())
		},
	}
)

func init() {
	ConfigCommand.AddCommand(
		ConfigPathCommand,
		ConfigLangCommand,
	)
}
