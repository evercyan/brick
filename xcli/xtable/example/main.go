package main

import (
	"fmt"

	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
)

var (
	structList = []struct {
		Name string `json:"name" table:"header-name"`
		Age  int    `json:"age"`
	}{
		{"Stark", 20},
		{"Lannister", 21},
	}

	numberList = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
)

func main() {
	fmt.Println("-------------------------------- border style")
	for s := xtable.Solid; s <= xtable.Dotted; s++ {
		xtable.New(structList).Style(s).Render()
	}

	fmt.Println("-------------------------------- other")
	fmt.Println("get header from  struct tag `table`:")
	xtable.New(structList).Render()
	fmt.Println("custom header:")
	xtable.New(structList).Header([]string{"col1", "col2"}).Border(true).Render()
	fmt.Println("show data border-bottom:")
	xtable.New(numberList).Border(true).Render()
	fmt.Println("with color:")
	xcolor.New(xtable.New(structList).Text()).Fg(xcolor.FgGreen).Render()
}
