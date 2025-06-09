package xtea

import (
	"testing"

	"github.com/evercyan/brick/xjson"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	options := []*SelectOption{
		{Label: "AAA", Desc: "I have ’em all over my house", Value: 1},
		{Label: "BBB", Desc: "It's good on toast", Value: "a"},
	}
	//options = []*SelectOption{
	//	{Label: "AAA"},
	//	{Label: "BBB"},
	//}
	option, err := Select("请选择选项", options)
	assert.Nil(t, err)
	xjson.Pretty(option, true)
}
