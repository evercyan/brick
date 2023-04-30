package xmerge

import (
	"image"
	"testing"

	"github.com/evercyan/brick/ximg"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	src := ximg.Read("../../logo.png")
	dst, err := Merge([]image.Image{src, src}, 1, 2)
	assert.Nil(t, err)
	assert.NotNil(t, dst)
}
