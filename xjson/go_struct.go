package xjson

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/evercyan/brick/xconvert"
)

// parser ...
type parser struct {
	Indent int
}

// Run ...
func (t *parser) Run(obj interface{}) string {
	return "type JSONToGO " + t.Parse(obj)
}

// Parse ...
func (t *parser) Parse(data interface{}) string {
	if v, ok := data.([]interface{}); ok {
		return t.ParseArray(v)
	}
	if v, ok := data.(map[string]interface{}); ok {
		return t.ParseObject(v)
	}
	return t.Type(data)
}

// ParseObject ...
func (t *parser) ParseObject(m map[string]interface{}) string {
	res := "struct {"
	res += "\n"
	lines := make([]string, 0)
	t.Indent += 1
	for k, v := range m {
		line := t.IndentString()
		line += xconvert.Title(xconvert.ToCamelCase(k)) + " "
		line += t.Parse(v)
		line += fmt.Sprintf(" `json:\"%s\"`", xconvert.ToSnakeCase(k))
		lines = append(lines, line)
	}
	t.Indent -= 1
	// 字段按字母顺序
	sort.Strings(lines)
	res += strings.Join(lines, "\n")
	res += "\n"
	res += t.IndentString() + "}"
	return res
}

func (t *parser) ParseArray(l []interface{}) string {
	// 默认无元素为 interface{}
	typeStr := "interface{}"
	for k, v := range l {
		vType := t.Type(v)
		if k == 0 {
			typeStr = vType
		}
		// 判断如果前后类型不一致则为 interface{}
		if typeStr != vType {
			typeStr = "interface{}"
			break
		}
	}
	res := "[]"
	if typeStr != "object" {
		res += typeStr
		return res
	}
	// 如果子元素类型为 struct, 需遍历列表中的每个对象, 将所有字段合并
	m := make(map[string]interface{})
	for _, v := range l {
		for kk, vv := range v.(map[string]interface{}) {
			if _, ok := m[kk]; !ok {
				m[kk] = vv
				continue
			}
			// 已存在该字段, 但前后子元素的变量类型不一致, 则赋值使解析时为 interface{}
			if t.Type(m[kk]) != t.Type(vv) {
				m[kk] = nil
			}
		}
	}
	res += t.ParseObject(m)
	return res
}

// Type ...
func (t *parser) Type(v interface{}) string {
	switch v.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return "int64"
	case float32, float64:
		// {"a": 1} 也会被解析成 float64
		// 处理时判断不包含 . 的则为 int64
		if !strings.Contains(fmt.Sprint(v), ".") {
			return "int64"
		}
		return "float64"
	case map[string]interface{}:
		return "object"
	case []interface{}:
		return "array"
	default:
		return "interface{}"
	}
}

// IndentString ...
func (t *parser) IndentString() string {
	return strings.Repeat("    ", t.Indent)
}

// ----------------------------------------------------------------

// ToGoStruct ...
func ToGoStruct(s string) (string, error) {
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return "", fmt.Errorf("无效的 JSON 字符串")
	}
	return new(parser).Run(v), nil
}
