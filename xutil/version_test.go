package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatVersion(t *testing.T) {
	assert.Equal(t, "1.2.3", FormatVersion("v1-2-3"))
	assert.Equal(t, "1.2a.3", FormatVersion("1.2a.3"))
}

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, 0, CompareVersion("v1.2.3", "V1-2-3"))
	assert.Equal(t, -1, CompareVersion("v1.2.3", "V1-10-3"))
	assert.Equal(t, 1, CompareVersion("v1.2.3", "V1-ab-3"))

	assert.Equal(t, -1, CompareVersion("", "v1.2.3"))
	assert.Equal(t, 1, CompareVersion("v1.2.3", ""))
	assert.Equal(t, 1, CompareVersion("v1.2.3", "v1.2"))
	assert.Equal(t, 0, CompareVersion("v1.2.0", "v1.2.a"))
	assert.Equal(t, -1, CompareVersion("v1.2", "v1.2.3"))
}
