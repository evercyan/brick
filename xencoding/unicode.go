package xencoding

import (
	"fmt"
	"strconv"
	"strings"
)

// UnicodeEncode ...
func UnicodeEncode(str string) string {
	quoted := strconv.QuoteToASCII(str)
	return quoted[1 : len(quoted)-1]
}

// UnicodeDecode ...
func UnicodeDecode(str string) string {
	res := ""
	for _, v := range strings.Split(str, "\\u") {
		if len(v) < 1 {
			continue
		}
		vv, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			return ""
		}
		res += fmt.Sprintf("%c", vv)
	}
	return res
}
