package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// GormCommand ...
	GormCommand = &cobra.Command{
		Use:   "gorm",
		Short: "Generate gorm struct",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("请输入 mysql 建表语句")
				return
			}
			res, err := app.GormStruct(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(res)
		},
	}
	// CommonCommand ...
	CommonCommand = &cobra.Command{
		Use:   "common",
		Short: "Generate common struct",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("请输入 mysql 建表语句")
				return
			}
			res, err := app.CommonStruct(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(res)
		},
	}
	// EnumCommand ...
	EnumCommand = &cobra.Command{
		Use:   "enum",
		Short: "Generate enum code",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("请输入待解析文本, e.g. AgeUnit:int8:年龄单位: 0, 岁; 1, 月; 2, 天;")
				return
			}
			res, err := app.Enum(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(res)
		},
	}
)
