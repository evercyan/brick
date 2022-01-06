package xutil

import (
	"fmt"
	"runtime"
)

// CallerLine ...
func CallerLine(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d(%s)", file, line, runtime.FuncForPC(pc).Name())
}

// CallerLines ...
func CallerLines(skip int) []string {
	res := make([]string, 0)
	for {
		s := CallerLine(skip)
		if s == "" {
			break
		}
		res = append(res, s)
		skip++
	}
	return res
}
