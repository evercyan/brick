package xloading

// Loading ...
type Loading interface {
	Style(Style) Loading   // 加载样式
	Symbol(Symbol) Loading // 结束字符
	Speed(int) Loading     // 加载速度
	Start() Loading        // 开始加载
	Success(...string)     // 加载成功
	Fail(...string)        // 加载失败
}
