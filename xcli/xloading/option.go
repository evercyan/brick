package xloading

// Style 加载样式
type Style int

// 样式
const (
	Style1 Style = iota
	Style2
	Style3
	Style4
	Style5
)

// Elements ...
func (t Style) Elements() []string {
	switch t {
	case Style2:
		return []string{"\\", "|", "/", "-"}
	case Style3:
		return []string{"🌕", "🌖", "🌗", "🌘", "🌑", "🌒", "🌓", "🌔"}
	case Style4:
		return []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}
	case Style5:
		return []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
	case Style1:
		fallthrough
	default:
		return []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	}
}

// ----------------------------------------------------------------

// Symbol 结束符号
type Symbol int

// 符号
const (
	Symbol1 Symbol = iota
	Symbol2
	Symbol3
	Symbol4
	Symbol5
)

// Elements ...
func (t Symbol) Elements() []string {
	switch t {
	case Symbol2:
		return []string{"✓", "✗"}
	case Symbol3:
		return []string{"✅", "❎"}
	case Symbol4:
		return []string{"👍", "👎"}
	case Symbol5:
		return []string{"💚", "💔"}
	case Symbol1:
		fallthrough
	default:
		return []string{"✔︎", "✘"}
	}
}
