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

// Success 输出成功
func Success(args ...string) {
	prefix(FgGreen, args...)
}

// Fail 输出失败
func Fail(args ...string) {
	prefix(FgRed, args...)
}

// Info 输出提示
func Info(args ...string) {
	prefix(FgYellow, args...)
}
