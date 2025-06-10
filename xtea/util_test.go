package xtea

import (
	"fmt"
	"testing"
)

func TestXtea(t *testing.T) {
	options := []string{
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
		"I have ’em all over my house",
		"It's good on toast",
	}
	option, err := Option("请选择选项", options)
	if err != nil {
		panic(err)
	}
	fmt.Println("option:", option)
	begin, err := Input("请输入开始日期", "2025-01-01")
	if err != nil {
		panic(err)
	}
	fmt.Println("begin:", begin)
	end, err := Input("请输入结束日期", "2025-01-01")
	if err != nil {
		panic(err)
	}
	fmt.Println("end:", end)
}
