package xcolor

import (
	"fmt"

	"github.com/evercyan/brick/xlodash"
)

// Output 处理提示前置
func Output(c Fg, texts ...interface{}) {
	text := ""
	if len(texts) > 1 {
		text += New(fmt.Sprint(texts[0])).Fg(c).Text() + " "
		texts = texts[1:]
	}
	text += New(xlodash.Map(texts, func(i int, v interface{}) string {
		return fmt.Sprint(v)
	})...).Fg(c).Text()
	fmt.Println(text)
}

// ----------------------------------------------------------------

// Success ...
func Success(args ...interface{}) {
	Output(FgGreen, args...)
}

// Info ...
func Info(args ...interface{}) {
	Output(FgCyan, args...)
}

// Warnning ...
func Warnning(args ...interface{}) {
	Output(FgYellow, args...)
}

// Danger ...
func Danger(args ...interface{}) {
	Output(FgRed, args...)
}
