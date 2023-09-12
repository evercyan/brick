package xalfred

import (
	"fmt"
)

// ----------------------------------------------------------------

// ScriptFilter 对应 Script Filter 组件返回
// https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type Response struct {
	Rerun          bool              `json:"rerun,omitempty"`          // 为 true 时再次执行(时间秒针变化)
	Skipknowledge  bool              `json:"skipknowledge,omitempty"`  // xxx
	Variables      map[string]string `json:"variables,omitempty"`      // 设置变量, 可被 items 使用
	Items          []*Item           `json:"items,omitempty"`          // 输出列表
	Alfredworkflow *Alfredworkflow   `json:"alfredworkflow,omitempty"` // 输出变量
}

// ----------------------------------------------------------------

// Alfredworkflow ...
type Alfredworkflow struct {
	Arg       string `json:"arg"`
	Config    *Var   `json:"config"`
	Variables *Var   `json:"variables"` // 通过 {var:xxx} 使用
}

// ----------------------------------------------------------------

// Item 对应 script filter 中的列表项
type Item struct {
	Uid          string      `json:"uid,omitempty"`          // 唯一标识
	Title        string      `json:"title"`                  // 主标题
	Subtitle     string      `json:"subtitle,omitempty"`     // 副标题
	Arg          string      `json:"arg,omitempty"`          // 传递的参数, 通过 {query} 访问
	Icon         *Icon       `json:"icon,omitempty"`         // 图标
	Valid        bool        `json:"valid,omitempty"`        // 是否有效, 默认 true, 无效时回车不执行任何操作
	Match        string      `json:"match,omitempty"`        // Filters Results
	Autocomplete string      `json:"autocomplete,omitempty"` // Tab 健自动补全
	Type         string      `json:"type,omitempty"`         // 类型, 默认 default(default/file/file:skipcheck)
	Mods         *Mod        `json:"mods,omitempty"`         // 控制按键回车时可变更 subtitle, icon, arg 等
	Action       interface{} `json:"action,omitempty"`       // 定义 Universal Action, 类型 object/array/string
	Text         *Text       `json:"text,omitempty"`         // 定义 cmd+c 复制和 cmd+l 显示时的文本
	Quicklookurl string      `json:"quicklookurl,omitempty"` // 快速预览, shift 或者 cmd+y, 无则使用 arg
	Variables    *Var        `json:"variables,omitempty"`    // 传递的变量, 通过 {var:xxx} 访问
}

// ----------------------------------------------------------------

// Mod ...
type Mod map[ModKey]*ModItem

// ModKey ...
type ModKey string

// ...
const (
	ModKeyCmd   = "cmd"   // ⌘
	ModKeyAlt   = "alt"   // ⌥
	ModKeyOpt   = "alt"   // ⌥
	ModKeyCtrl  = "ctrl"  // ^
	ModKeyShift = "shift" // ⇧
	ModKeyFn    = "fn"    // fn
)

// Plus 支持组合 cmd + shift
func (t ModKey) Plus(k ModKey) ModKey {
	return ModKey(fmt.Sprintf("%s+%s", t, k))
}

// ModItem ...
type ModItem struct {
	Key       string   `json:"key,omitempty"`
	Subtitle  string   `json:"subtitle,omitempty"`
	Icon      *Icon    `json:"icon,omitempty"`
	Valid     bool     `json:"valid,omitempty"`
	Arg       []string `json:"arg,omitempty"`
	Variables *Var     `json:"variables,omitempty"`
}

// ----------------------------------------------------------------

// Text ...
type Text struct {
	Copy      string `json:"copy,omitempty"`      // cmd+c 复制
	Largetype string `json:"largetype,omitempty"` // cmd+l 显示
}

// ----------------------------------------------------------------

// IconType 图标类型
type IconType string

// ...
const (
	// 图片, 值为对应图片路径
	IconTypeImage IconType = ""
	// 文件图标, "/Applications/Safari.app" 即显示对应应用的图片
	IconTypeFileIcon IconType = "fileicon"
	// 文件类型, "public.folder" "com.apple.rtfd"
	IconTypeFileType IconType = "filetype"
)

// ----------------------------------------------------------------

// Icon ...
type Icon struct {
	Type IconType `json:"type,omitempty"`
	Path string   `json:"path"`
}

// ----------------------------------------------------------------

// Var ...
type Var map[string]interface{}
