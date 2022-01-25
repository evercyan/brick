package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	assert.Equal(t, 0, RandNumber(2, 1))
	assert.LessOrEqual(t, 1, RandNumber(1, 10))
	assert.GreaterOrEqual(t, 10, RandNumber(1, 10))

	assert.LessOrEqual(t, -10, RandNumber(-10, 10))
	assert.GreaterOrEqual(t, 10, RandNumber(-10, 10))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Range(2, 1))
	assert.Equal(t, []int{1, 2}, Range(1, 2))
}

func TestRandString(t *testing.T) {
	assert.Equal(t, 6, len(RandString(6)))
	assert.NotEqual(t, RandString(6), RandString(6))
}

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

func TestPretty(t *testing.T) {
	data := map[string]string{
		"hello": "world",
	}
	assert.NotEmpty(t, Pretty(data))
}

func TestLen(t *testing.T) {
	s := "hello 世界"
	assert.Equal(t, 12, len(s))
	assert.Equal(t, 8, Len(s))
}
