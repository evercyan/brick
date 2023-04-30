package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	Setenv("hello", "world")

	assert.Equal(t, "world", Getenv("hello"))
	assert.Equal(t, "", Getenv("hello1"))
	assert.Equal(t, "wooo", Getenv("hello1", "wooo"))

	envMap := GetenvMap()
	assert.Equal(t, "world", envMap["hello"])
}

func TestPlatform(t *testing.T) {
	if IsMac() {
		assert.True(t, IsMac())
		assert.False(t, IsWin())
		assert.False(t, IsLinux())
	} else if IsWin() {
		assert.False(t, IsMac())
		assert.True(t, IsWin())
		assert.False(t, IsLinux())
	} else if IsLinux() {
		assert.False(t, IsMac())
		assert.False(t, IsWin())
		assert.True(t, IsLinux())
	}
}
