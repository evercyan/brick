package xutil

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

// RandNumber ...
func RandNumber(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max+1-min)
}

// RandString ...
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// Range ...
func Range(min, max int) []int {
	res := make([]int, 0)
	if min > max {
		min, max = max, min
	}
	for i := min; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

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

// Pretty ...
func Pretty(v interface{}) string {
	out, _ := json.MarshalIndent(v, "", "    ")
	return string(out)
}

// Len ..
func Len(s string) int {
	return utf8.RuneCountInString(s)
}
