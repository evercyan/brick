package xcolor

import (
	"fmt"

	"github.com/evercyan/brick/xlodash"
)

// prefix 处理提示前置
func prefix(c Fg, texts ...interface{}) {
	text := ""
	if len(texts) > 1 {
		text += New(fmt.Sprint(texts[0])).Fg(FgYellow).Text() + " "
		texts = texts[1:]
	}
	text += New(xlodash.Map(texts, func(i int, v interface{}) string {
		return fmt.Sprint(v)
	})...).Fg(c).Text()
	fmt.Println(text)
}

// Success 输出成功
func Success(args ...interface{}) {
	prefix(FgGreen, args...)
}

// Fail 输出失败
func Fail(args ...interface{}) {
	prefix(FgRed, args...)
}

// Info 输出提示
func Info(args ...interface{}) {
	prefix(FgYellow, args...)
}
