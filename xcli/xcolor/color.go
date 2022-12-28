package xcolor

// Package xcolor 渲染终端文本颜色背景色和显示样式
//  xcolor.New("text").Fg(xcolor.FgGreen).Render()

import (
	"fmt"
	"strings"
)

type color struct {
	text string
	fg   Fg
	bg   Bg
	sty  Sty
}

// New ...
func New(texts ...string) Color {
	return &color{
		text: strings.Join(texts, " "),
	}
}

// ----------------------------------------------------------------

// Fg 设置前景色
func (t *color) Fg(v Fg) Color {
	if v.String() != "" {
		t.fg = v
	}
	return t
}

// Bg 设置背景色
func (t *color) Bg(v Bg) Color {
	if v.String() != "" {
		t.bg = v
	}
	return t
}

// Sty 设置样式
func (t *color) Sty(v Sty) Color {
	if v.String() != "" {
		t.sty = v
	}
	return t
}

// Text 返回渲染后文本
func (t *color) Text() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, t.sty, t.bg, t.fg, t.text, 0x1B)
}

// Render 输出渲染后文本
func (t *color) Render() {
	fmt.Println(t.Text())
}
