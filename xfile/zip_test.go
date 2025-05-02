package xfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteZip(t *testing.T) {
	err := WriteZip("./test.zip", []string{
		"./dir.go",
		"../xlog",
	}, WithZipPassword("123456"), WithZipKeepLevel(true))
	assert.Nil(t, err)
	os.Remove("./test.zip")
}
