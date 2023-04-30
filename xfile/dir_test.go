package xfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDir(t *testing.T) {
	assert.NotEmpty(t, GetCurrentDir())
	assert.NotEmpty(t, GetHomeDir())
	_, err := GetConfigDir("kepler")
	assert.Nil(t, err)
}

func TestList(t *testing.T) {
	assert.Less(t, 0, len(ListFiles("./", "")))
	assert.Less(t, 0, len(ListFiles("../", "is", true)))
	assert.Equal(t, 0, len(ListFiles("./aaa", "")))
	assert.Equal(t, 0, len(ListDirs("./")))
	assert.Less(t, 0, len(ListDirs("../", true)))
}
