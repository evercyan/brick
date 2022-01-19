package xversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	assert.Equal(t, "1.2.3", Format("v1-2-3"))
	assert.Equal(t, "1.2a.3", Format("1.2a.3"))
}

func TestCompare(t *testing.T) {
	assert.Equal(t, Equal, Compare("v1.2.3", "V1-2-3"))
	assert.Equal(t, Less, Compare("v1.2.3", "V1-10-3"))
	assert.Equal(t, Greater, Compare("v1.2.3", "V1-ab-3"))

	assert.Equal(t, Less, Compare("", "v1.2.3"))
	assert.Equal(t, Greater, Compare("v1.2.3", ""))
	assert.Equal(t, Greater, Compare("v1.2.3", "v1.2"))
	assert.Equal(t, Equal, Compare("v1.2.0", "v1.2.a"))
	assert.Equal(t, Less, Compare("v1.2", "v1.2.3"))
}
