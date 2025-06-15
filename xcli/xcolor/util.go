package xcolor

import (
	"fmt"

	"github.com/evercyan/brick/xlodash"
)

// ...
const (
	F = "✘"
	S = "✔︎"
	I = "➤"
	L = "────────────────────────────────────────────────────────────────────────────────────────────────"
)

// ----------------------------------------------------------------

// Text 处理提示前置
func Text(c Fg, texts ...interface{}) string {
	text := ""
	if len(texts) > 1 {
		text += New(fmt.Sprint(texts[0])).Fg(c).Text() + " "
		texts = texts[1:]
	}
	text += New(xlodash.Map(texts, func(i int, v interface{}) string {
		return fmt.Sprint(v)
	})...).Fg(c).Text()
	return text
}

// Output 处理提示前置
func Output(c Fg, texts ...interface{}) {
	fmt.Println(Text(c, texts...))
}

// ----------------------------------------------------------------

// Success ...
func Success(args ...interface{}) {
	args = append([]interface{}{S}, args...)
	Output(FgGreen, args...)
}

// Info ...
func Info(args ...interface{}) {
	args = append([]interface{}{I}, args...)
	Output(FgCyan, args...)
}

// Warnning ...
func Warnning(args ...interface{}) {
	args = append([]interface{}{I}, args...)
	Output(FgYellow, args...)
}

// Danger ...
func Danger(args ...interface{}) {
	args = append([]interface{}{F}, args...)
	Output(FgRed, args...)
}

// Success ...
func Successf(tpl string, args ...interface{}) {
	Success(fmt.Sprintf(tpl, args...))
}

// Info ...
func Infof(tpl string, args ...interface{}) {
	Info(fmt.Sprintf(tpl, args...))
}

// Warnning ...
func Warnningf(tpl string, args ...interface{}) {
	Warnning(fmt.Sprintf(tpl, args...))
}

// Danger ...
func Dangerf(tpl string, args ...interface{}) {
	Danger(fmt.Sprintf(tpl, args...))
}

// Successt ...
func Successt(args ...interface{}) string {
	args = append([]interface{}{S}, args...)
	return Text(FgGreen, args...)
}

// Infot ...
func Infot(args ...interface{}) string {
	args = append([]interface{}{I}, args...)
	return Text(FgCyan, args...)
}

// Warnningt ...
func Warnningt(args ...interface{}) string {
	args = append([]interface{}{I}, args...)
	return Text(FgYellow, args...)
}

// Dangert ...
func Dangert(args ...interface{}) string {
	args = append([]interface{}{F}, args...)
	return Text(FgRed, args...)
}
