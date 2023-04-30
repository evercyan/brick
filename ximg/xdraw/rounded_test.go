package xdraw

import (
	"testing"

	"github.com/evercyan/brick/ximg"
	"github.com/stretchr/testify/assert"
)

func TestRounded(t *testing.T) {
	src := ximg.Read("../../logo.png")
	dst := Rounded(src, 0.25)
	assert.NotNil(t, dst)

	assert.NotNil(t, Rounded(src, 10))
	assert.NotNil(t, Circle(src))
}

func TestBorder(t *testing.T) {
	src := ximg.Read("../../logo.png")
	dst := Border(src, 100)
	assert.NotNil(t, dst)

	assert.NotNil(t, Border(src, -1))
	assert.NotNil(t, Border(src, 10, "#ff0000"))
}
