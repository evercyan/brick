package xlodash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3}, 1))
	assert.True(t, Contains([]string{"a", "b"}, "a"))
	assert.False(t, Contains([]string{"a", "b"}, "c"))
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, 0, IndexOf([]int{1, 2, 3}, 1))
	assert.Equal(t, -1, IndexOf([]int{1, 2, 3}, 10))
	assert.Equal(t, 2, LastIndexOf([]int{1, 2, 1}, 1))
	assert.Equal(t, -1, LastIndexOf([]int{1, 2, 1}, 10))
}

func TestFind(t *testing.T) {
	l1 := []string{"a", "b"}
	item, index := Find(l1, func(index int, item string) bool {
		return item == "b"
	})
	assert.Equal(t, index, 1)
	assert.Equal(t, item, "b")
}
