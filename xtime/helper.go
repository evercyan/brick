package xtime

import (
	"strings"
)

var (
	layoutMap = map[string]string{
		"y": "2006",
		"m": "01",
		"d": "02",
		"h": "15",
		"i": "04",
		"s": "05",
	}
)

// replaceLayout ...
func replaceLayout(layout string) string {
	layout = strings.ToLower(layout)
	for k, v := range layoutMap {
		layout = strings.ReplaceAll(layout, k, v)
	}
	return layout
}
