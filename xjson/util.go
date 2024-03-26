package xjson

import (
	"encoding/json"
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
	if vv, ok := v.(string); ok {
		if vvv, err := strconv.Unquote(vv); err == nil {
			vv = vvv
		}
		var raw interface{}
		if err := json.Unmarshal([]byte(vv), &raw); err == nil {
			v = raw
		}
	}
	return v
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

// FilterPrefix 过滤转义字符
func FilterTransfer(s string) string {
	s = strings.ReplaceAll(s, `\"`, `"`)
	s = strings.ReplaceAll(s, `\\`, `\`)
	return s
}

// Format 格式化成可用 JSON 字符串
func Format(s string) string {
	filters := []func(string) string{FilterPrefix, FilterSuffix, FilterTransfer}
	for _, fn := range filters {
		if xtype.IsJSONString(s) {
			break
		}
		s = fn(s)
	}
	return s
}
