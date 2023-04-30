package xclip

import (
	"testing"

	"github.com/evercyan/brick/xfile"
	"github.com/stretchr/testify/assert"
)

func TestXclip(t *testing.T) {
	Write("hello")
	Read()
	assert.NotNil(t, ReadImage(xfile.Temp()))
	WriteImage("../../logo.png")
	assert.Nil(t, ReadImage(xfile.Temp()))
}
