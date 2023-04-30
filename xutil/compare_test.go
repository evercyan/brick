package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	assert.Equal(t, 0, Compare(1, 1))
	assert.Equal(t, 1, Compare(2, 1))
	assert.Equal(t, -1, Compare(1, 2))
	assert.Equal(t, 0, Compare("1.0", 1))
	assert.Equal(t, 0, Compare("a", "a"))
	assert.Equal(t, 0, Compare(3, 3.0))
	assert.Equal(t, 0, Compare(1, true))
	assert.Equal(t, 0, Compare(0, false))
	assert.Equal(t, 0, Compare(nil, nil))
	assert.Equal(t, 0, Compare(struct{}{}, struct{}{}))
}
