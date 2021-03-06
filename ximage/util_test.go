package ximage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.Equal(t, "png", Type("../logo.png"))
	assert.Empty(t, Type("../xfile"))
	assert.Empty(t, Type("../README.md"))
}

func TestSize(t *testing.T) {
	w, h := Size("../logo.png")
	assert.Equal(t, 512, w)
	assert.Equal(t, 512, h)

	w1, h1 := Size("../README.md")
	assert.Empty(t, w1)
	assert.Empty(t, h1)
}
