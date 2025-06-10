package xtea

import (
	"testing"

	"github.com/evercyan/brick/xjson"
	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	options := []string{"I have ’em all over my house", "It's good on toast"}
	option, err := Option("请选择选项", options)
	assert.Nil(t, err)
	xjson.Pretty(option, true)
}
