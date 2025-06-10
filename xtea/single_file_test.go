package xtea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleFile(t *testing.T) {
	fpath, err := SingleFile("../", ".go")
	assert.Nil(t, err)
	fmt.Println(fpath)
}
