package xutil

import (
	"math"
	"strings"
	"unicode/utf8"

	"github.com/evercyan/brick/xtype"
	"github.com/mozillazg/go-pinyin"
)

// If 三目运算
func If(cond bool, val1, val2 interface{}) interface{} {
	if cond {
		return val1
	}
	return val2
}

// Replace ...
func Replace(s string, replace map[string]string) string {
	for k, v := range replace {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

// Len ..
func Len(s string) int {
	return utf8.RuneCountInString(s)
}

// Default ...
func Default[V comparable](v V, de V) V {
	if xtype.IsZero(v) {
		return de
	}
	return v
}

// Round ...
func Round(x float64, length int) float64 {
	factor := math.Pow10(length)
	return math.Round(x*factor) / factor
}

// Ceil ...
func Ceil(x float64, length int) float64 {
	factor := math.Pow10(length)
	return math.Ceil(x*factor) / factor
}

// Abbr ...
func Abbr(name string) string {
	items := pinyin.LazyConvert(name, nil)
	if len(items) == 0 {
		return ""
	}
	abbrs := make([]string, 0)
	for _, item := range items {
		if len(item) == 0 {
			continue
		}
		abbrs = append(abbrs, string(item[0]))
	}
	return strings.Join(abbrs, "")
}
