package ximg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	img, err := Compress("../logo.png", 10)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.NotNil(t, img)
}
