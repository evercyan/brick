package xtable

// Table ...
type Table interface {
	Style(int) Table       // 设置边界样式
	Border(bool) Table     // 设置是否显示内容区下边界
	Header([]string) Table // 设置表格头部数据
	Text() string          // 返回渲染后文本
	Render()               // 输出渲染后文本
}
