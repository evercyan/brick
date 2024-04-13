package ximg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResize(t *testing.T) {
	src := Read("../../logo.png")
	dst := Resize(src, 100, 100)
	assert.NotNil(t, dst)
}
