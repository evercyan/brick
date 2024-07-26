package xjson

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	"github.com/evercyan/brick/xencoding"
	"github.com/evercyan/brick/xtype"
)

// ...
var (
	Encode = xencoding.JSONEncode
	Decode = xencoding.JSONDecode
)

// format ...
func format(v interface{}) interface{} {
	vv, ok := v.(string)
	if !ok {
		return v
	}
	var vvv interface{}
	if err := json.Unmarshal([]byte(vv), &vvv); err != nil {
		return v
	}
	return vvv
}

// Pretty ...
func Pretty(v interface{}) string {
	b, _ := json.MarshalIndent(format(v), "", strings.Repeat(" ", 4))
	return string(b)
}

// Minify ...
func Minify(v interface{}) string {
	return Encode(format(v))
}

// FilterPrefix 过滤前置字符直到 { 或 [
func FilterPrefix(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' {
			s = s[i:]
			break
		}
	}
	return s
}

// FilterSuffix 过滤后置字符直到 } 或 ]
func FilterSuffix(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '}' || s[i] == ']' {
			s = s[:i+1]
			break
		}
	}
	return s
}

// FilterUnquote 过滤转义字符
func FilterUnquote(s string) string {
	if v, err := strconv.Unquote(s); err == nil {
		return v
	}
	s = strings.ReplaceAll(s, `\"`, `"`)
	s = strings.ReplaceAll(s, `\\`, `\`)
	return s
}

// Format 格式化成可用 JSON 字符串
func Format(s string) string {
	filters := []func(string) string{FilterPrefix, FilterSuffix, FilterUnquote}
	for _, fn := range filters {
		if xtype.IsJSONString(s) {
			break
		}
		s = fn(s)
	}
	return s
}

// Sort ...
func Sort(s string) string {
	var list []interface{}
	if err := json.Unmarshal([]byte(s), &list); err != nil {
		return s
	}
	if len(list) == 0 {
		return s
	}
	switch list[0].(type) {
	case float64:
		sort.Slice(list, func(i, j int) bool {
			return xtype.ToFloat64(list[i]) < xtype.ToFloat64(list[j])
		})
		return Encode(list)
	case string:
		isNumber := xtype.ToInt64(list[0]) > 0
		sort.Slice(list, func(i, j int) bool {
			if isNumber {
				return xtype.ToFloat64(list[i]) < xtype.ToFloat64(list[j])
			}
			return xtype.ToString(list[i]) < xtype.ToString(list[j])
		})
		return Encode(list)
	default:
		return s
	}
}
