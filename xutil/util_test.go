package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	x, y := 5, 6
	assert.Equal(t, y, If(false, x, y))
	assert.Equal(t, x, If(true, x, y))
}

func TestReplace(t *testing.T) {
	assert.Equal(t, "ifllo", Replace("hello", map[string]string{
		"h": "i",
		"e": "f",
	}))
}

func TestLen(t *testing.T) {
	s := "hello 世界"
	assert.Equal(t, 12, len(s))
	assert.Equal(t, 8, Len(s))
}
