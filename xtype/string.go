package xtype

import (
	"strings"

	"github.com/evercyan/brick/xlodash"
)

// String2List ...
func String2List(str string, sep string, uniques ...bool) []string {
	list := make([]string, 0)
	if str == "" {
		return list
	}
	unique := xlodash.First(uniques, false)
	m := make(map[string]struct{})
	for _, v := range strings.Split(str, sep) {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if unique {
			if _, ok := m[v]; ok {
				continue
			}
			m[v] = struct{}{}
		}
		list = append(list, v)
	}
	return list
}
