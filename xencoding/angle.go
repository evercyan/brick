package xencoding

import (
	"strings"
)

var (
	angleMap = map[string]string{
		"，": ",",
		"。": ".",
		"！": "!",
		"？": "?",
		"、": ",",
		"；": ";",
		"：": ":",
		"～": "~",
		"‘": "'",
		"“": `"`,
		"”": `"`,
		"（": "(",
		"）": ")",
		"《": "<",
		"》": ">",
		"〈": "<",
		"〉": ">",
		"【": "[",
		"】": "]",
	}
)

// Ord ...
func Ord(s string) int {
	return int([]rune(s)[0])
}

// Chr ...
func Chr(s int) string {
	return string(rune(s))
}

// FullAngle ...
func FullAngle(s string) string {
	for k, v := range angleMap {
		s = strings.ReplaceAll(s, v, k)
	}
	return s
}

// HalfAngle ...
func HalfAngle(s string) string {
	for k, v := range angleMap {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
