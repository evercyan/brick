package xpdf

// Line ...
type Line struct {
	Type Type
	Text string
	Size float64
}

// ----------------------------------------------------------------

// Type 类型
type Type int

const (
	TypeText Type = iota
	TypeImage
	TypeDivider
	TypePage
)

func (t Type) String() string {
	switch t {
	case TypeText:
		return "文本"
	case TypeImage:
		return "图片"
	case TypeDivider:
		return "分隔线"
	case TypePage:
		return "分页"
	default:
		return ""
	}
}
