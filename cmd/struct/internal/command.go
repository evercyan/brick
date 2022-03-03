package internal

import (
	"fmt"
	"os"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xfile"
	"github.com/spf13/cobra"
)

var (
	// GormCommand ...
	GormCommand = &cobra.Command{
		Use:   "gorm",
		Short: "生成 Gorm Struct",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail("Error:", "请输入 mysql 建表语句")
				return
			}
			res, err := app.GormStruct(args[0])
			if err != nil {
				xcolor.Fail("Error:", err.Error())
				return
			}
			fmt.Println(res)
		},
	}
	// CommonCommand ...
	CommonCommand = &cobra.Command{
		Use:   "common",
		Short: "生成普通 Struct",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail("Error:", "请输入 mysql 建表语句")
				return
			}
			res, err := app.CommonStruct(args[0])
			if err != nil {
				xcolor.Fail("Error:", err.Error())
				return
			}
			fmt.Println(res)
		},
	}
	// EnumCommand ...
	EnumCommand = &cobra.Command{
		Use:   "enum",
		Short: "生成枚举代码",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail("Error:", "请输入待解析文本, e.g. struct enum AgeUnit:int8:年龄单位: 0, 岁; 1, 月; 2, 天;")
				return
			}
			res, err := app.Enum(args[0])
			if err != nil {
				xcolor.Fail("Error:", err.Error())
				return
			}
			fmt.Println(res)
		},
	}
	// SqlCommand ...
	SqlCommand = &cobra.Command{
		Use:   "sql",
		Short: "解析 sql 文件批量生成 Gorm Struct 文件",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				xcolor.Fail("Error:", "缺少 sql 文件路径, e.g. struct sql ~/Downloads/demo.sql")
				return
			}
			filepath := args[0]
			if !xfile.IsExist(filepath) || !xfile.IsFile(filepath) {
				xcolor.Fail("Error:", "无效的 sql 文件")
				return
			}

			dstDir := fmt.Sprintf("%s-output", filepath)
			if !xfile.IsExist(dstDir) {
				os.MkdirAll(dstDir, os.ModePerm)
			}
			app.Sql(filepath, dstDir)
		},
	}
)
