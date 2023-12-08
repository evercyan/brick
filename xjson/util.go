package xjson

import (
	"encoding/json"
	"regexp"
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

// Pretty ...
func Pretty(v interface{}) string {
	if vv, ok := v.(string); ok {
		if vvv, err := strconv.Unquote(vv); err == nil {
			vv = vvv
		}
		var raw json.RawMessage
		if err := json.Unmarshal([]byte(vv), &raw); err == nil {
			v = raw
		}
	}
	b, _ := json.MarshalIndent(v, "", strings.Repeat(" ", 4))
	return string(b)
}

// Minify ...
func Minify(v interface{}) string {
	if vv, ok := v.(string); ok {
		if vvv, err := strconv.Unquote(vv); err == nil {
			vv = vvv
		}
		var raw json.RawMessage
		if err := json.Unmarshal([]byte(vv), &raw); err == nil {
			v = raw
		}
	}
	return Encode(v)
}

// Format 格式化成可用 JSON 字符串
func Format(v string) string {
	if xtype.IsJSONString(v) {
		return v
	}
	if regexp.MustCompile(`[^\\]"`).MatchString(v) {
		return v
	}
	v = strings.ReplaceAll(v, `\"`, `"`)
	v = strings.ReplaceAll(v, `\\`, `\`)
	return v
}
