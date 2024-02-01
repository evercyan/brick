package xalfred

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/evercyan/brick/xlodash"
)

// NewIcon ...
func NewIcon(path string, types ...IconType) *Icon {
	t := IconTypeImage
	if len(types) > 0 {
		t = types[0]
	}
	return &Icon{
		Type: t,
		Path: path,
	}
}

// NewItem ...
func NewItem(title, subtitle string, args ...string) *Item {
	return &Item{
		Title:    title,
		Subtitle: subtitle,
		Arg:      xlodash.First(args, subtitle),
	}
}

// GetSysIconPath 获取系统图片路径
func GetSysIconPath(name string) string {
	return fmt.Sprintf("/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/%s", name)
}

// ----------------------------------------------------------------

// List 列表, 适用于 Script Filter 列表输出
func List(items []*Item) {
	Output(&Response{
		Items: items,
	})
}

// Row 错误或提示, 适用于 Script Filter 单行输出
func Row(title, subtitle string) {
	Output(&Response{
		Items: []*Item{{
			Title:    title,
			Subtitle: subtitle,
		}},
	})
}

// Error 错误
var Error = Row

// Notice 通知, 适用于 Post Notification
func Notice(title, content string) {
	Output(&Response{
		Alfredworkflow: &Alfredworkflow{
			Variables: &Var{
				"title":   title,
				"content": content,
			},
		},
	})
}

// Dict 变量, 适用于 Run Script 输出
func Dict(v *Var) {
	Output(&Response{
		Alfredworkflow: &Alfredworkflow{
			Variables: v,
		},
	})
}

// ----------------------------------------------------------------

// Output ...
func Output(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
	os.Exit(0)
}
