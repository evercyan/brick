package xcolor

// Sty 样式枚举
type Sty int

// 样式
const (
	StyNone         Sty = iota // 默认
	StyBold                    // 加粗
	StyFaint                   // 弱化
	StyItalic                  // 斜体
	StyUnderline               // 下划
	StyBlinkSlow               // 慢闪
	StyBlinkRapid              // 快闪
	StyReverseVideo            // 反转
	StyConcealed               // 隐藏
	StyCrossedOut              // 中划
)

// String 样式描述
func (t Sty) String() string {
	switch t {
	case StyNone:
		return "默认"
	case StyBold:
		return "加粗"
	case StyFaint:
		return "弱化"
	case StyItalic:
		return "斜体"
	case StyUnderline:
		return "下划"
	case StyBlinkSlow:
		return "慢闪"
	case StyBlinkRapid:
		return "快闪"
	case StyReverseVideo:
		return "反转"
	case StyConcealed:
		return "隐藏"
	case StyCrossedOut:
		return "中划"
	default:
		return ""
	}
}

// ----------------------------------------------------------------

// Fg 前景色枚举 (Basic 30, Hi-Intensity 90)
type Fg int

// 前景色
const (
	FgNone    Fg = iota + 89 // 默认
	FgBlack                  // 黑色
	FgRed                    // 红色
	FgGreen                  // 绿色
	FgYellow                 // 黄色
	FgBlue                   // 蓝色
	FgMagenta                // 洋红
	FgCyan                   // 青色
	FgWhite                  // 白色
)

// String 前景色描述
func (t Fg) String() string {
	switch t {
	case FgNone:
		return "默认"
	case FgBlack:
		return "黑色"
	case FgRed:
		return "红色"
	case FgGreen:
		return "绿色"
	case FgYellow:
		return "黄色"
	case FgBlue:
		return "蓝色"
	case FgMagenta:
		return "洋红"
	case FgCyan:
		return "青色"
	case FgWhite:
		return "白色"
	default:
		return ""
	}
}

// ----------------------------------------------------------------

// Bg 背景色枚举 (Basic 40, Hi-Intensity 100)
type Bg int

// 背景色
const (
	BgNone    Bg = iota + 99 // 默认背景
	BgBlack                  // 黑色背景
	BgRed                    // 红色背景
	BgGreen                  // 绿色背景
	BgYellow                 // 黄色背景
	BgBlue                   // 蓝色背景
	BgMagenta                // 洋红背景
	BgCyan                   // 青色背景
	BgWhite                  // 白色背景
)

// String 背景色描述
func (t Bg) String() string {
	switch t {
	case BgNone:
		return "默认背景"
	case BgBlack:
		return "黑色背景"
	case BgRed:
		return "红色背景"
	case BgGreen:
		return "绿色背景"
	case BgYellow:
		return "黄色背景"
	case BgBlue:
		return "蓝色背景"
	case BgMagenta:
		return "洋红背景"
	case BgCyan:
		return "青色背景"
	case BgWhite:
		return "白色背景"
	default:
		return ""
	}
}
