package xcolor

// Color ...
type Color interface {
	Fg(Fg) Color   // 设置前景色
	Bg(Bg) Color   // 设置背景色
	Sty(Sty) Color // 设置样式
	Text() string  // 返回渲染后文本
	Render()       // 输出渲染后文本
}
