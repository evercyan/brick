package xtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	Name string `json:"name" table:"名称"`
	Age  int    `json:"age"`
}

var (
	structList = []user{
		{"Stark", 20},
		{"Lannister", 21},
	}

	structPtrList = []*user{
		{
			"Stark",
			20,
		},
		{
			"Lannister",
			21,
		},
	}

	numberList = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	stringList = [][]string{
		{"a", "bb", "ccc"},
		{"dddd", "eeeee", "ffffff"},
	}
)

func TestTable(t *testing.T) {
	assert.NotEmpty(t, New(structList).Style(Dashed).Border(true).Text())
	assert.NotEmpty(t, New(structList).Style(100).Text())
	assert.NotEmpty(t, New(structList).Header([]string{"Cooooooooooooool1", "Col2"}).Text())
	assert.NotEmpty(t, New(numberList).Header([]string{"Cooooooooooooool1", "Col2", "Col3"}).Text())
	assert.NotEmpty(t, New(stringList).Text())
	New(numberList).Render()
}

func TestTableCoverage(t *testing.T) {
	assert.NotEmpty(t, New(numberList).Style(100).Text())
	assert.NotNil(t, New(1).Text())
	assert.NotNil(t, New(structPtrList).Text())
	assert.NotNil(t, New([]int{1}).Text())
	assert.NotNil(t, New([][]map[string]string{
		{
			{
				"a": "a",
			},
		},
	}).Text())
	assert.NotNil(t, New([][]int{}).Text())
	assert.NotNil(t, New(structList).Header([]string{"c1"}).Text())
	assert.NotNil(t, New(numberList).Header([]string{"c1"}).Text())
}

func TestTableMarkdown(t *testing.T) {
	New(structList).Style(Dashed).Border(true).Render()
	New(structList).Style(Markdown).Border(true).Render()
}
