package main

import (
	"fmt"

	"github.com/evercyan/brick/xcli/xcolor"
)

func main() {
	fmt.Println("")
	// 显示样式
	for sty := xcolor.StyNone; sty <= xcolor.StyCrossedOut; sty++ {
		// 背景颜色
		for bg := xcolor.BgNone; bg <= xcolor.BgWhite; bg++ {
			// 字体颜色
			for fg := xcolor.FgNone; fg <= xcolor.FgWhite; fg++ {
				fmt.Printf(xcolor.New(
					sty.String(),
					bg.String(),
					fg.String(),
				).Sty(sty).Fg(fg).Bg(bg).Text() + " ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
