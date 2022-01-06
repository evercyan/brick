package xcolor

import (
	"fmt"
)

// prefix 处理提示前置
func prefix(c Fg, texts ...string) {
	text := ""
	if len(texts) > 1 {
		text += New(texts[0]).Fg(FgYellow).Text() + " "
		texts = texts[1:]
	}
	text += New(texts...).Fg(c).Text()
	fmt.Println(text)
}

// Success 输出成功提示
func Success(texts ...string) {
	prefix(FgGreen, texts...)
}

// Fail 输出失败文字
func Fail(texts ...string) {
	prefix(FgRed, texts...)
}
