package xtea

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	value, err := Input("请输入选项", "2025-06-09")
	assert.Nil(t, err)
	fmt.Println("TestInput value:", value)
}
