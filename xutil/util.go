package xutil

import (
	"strings"
	"unicode/utf8"
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
