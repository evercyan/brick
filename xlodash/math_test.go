package xlodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, "a", Min([]string{"a", "b"}...))
	assert.Equal(t, 1, Min([]int{2, 3, 1}...))

	assert.Equal(t, 0, Min[int]())
	assert.Equal(t, "", Min[string]())
}

func TestMax(t *testing.T) {
	assert.Equal(t, "b", Max([]string{"a", "b"}...))
	assert.Equal(t, 3, Max([]int{2, 3, 1}...))

	assert.Equal(t, 0, Max[int]())
	assert.Equal(t, "", Max[string]())
}
