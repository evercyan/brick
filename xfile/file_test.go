package xfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemp(t *testing.T) {
	assert.NotEmpty(t, Temp())
	assert.NotEmpty(t, Temp("abc"))
	assert.NotEmpty(t, Temp(".png"))
}

func TestSize(t *testing.T) {
	assert.Equal(t, int64(0), Size("./file_t.go"))
	assert.LessOrEqual(t, int64(0), Size("./file.go"))
}

func TestSizeText(t *testing.T) {
	assert.Equal(t, "0.00B", SizeText(0))
	assert.Equal(t, "1023.00B", SizeText(1023))
	assert.Equal(t, "1.00KB", SizeText(1024))
	assert.Equal(t, "1.65KB", SizeText(1024+666))
	assert.Equal(t, "1.65MB", SizeText((1024+666)*1024))
	assert.Equal(t, "1.65GB", SizeText((1024+666)*1024*1024))
	assert.Equal(t, "1.00TB", SizeText(1*1024*1024*1024*1024))
}

func TestRead(t *testing.T) {
	assert.NotEmpty(t, Read("./file.go"))
	assert.Empty(t, Read("./file_t.go"))

	dst := Temp()
	assert.Nil(t, Write(dst, "hello"))
	assert.Nil(t, Copy(dst, Temp()))
}

func TestLine(t *testing.T) {
	filepath := "../LICENSE"

	assert.Equal(t, 0, LineCount(filepath+"s"))
	assert.Equal(t, map[int]string(nil), LineContent(filepath+"s"))

	assert.Equal(t, 21, LineCount(filepath))
	assert.Equal(t, map[int]string{
		1: "MIT License",
		2: "",
		3: "Copyright (c) 2021 Ryan",
	}, LineContent(filepath, 1, 2, 3))
}

func TestMineType(t *testing.T) {
	assert.Equal(t, "image/png", MineType("../logo.png"))
	assert.Equal(t, "", MineType("../logo.png1111"))
	assert.Equal(t, "text/plain; charset=utf-8", MineType("../LICENSE"))
}

func TestJSON(t *testing.T) {
	type xjson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	data := &xjson{
		Name: "abc",
		Age:  1,
	}
	filepath := Temp("file.json")
	defer os.Remove(filepath)

	assert.Nil(t, WriteJSON(filepath, data))
	assert.Nil(t, WriteJSON(filepath, data, true))
	assert.True(t, IsExist(filepath))

	d := &xjson{}
	assert.Nil(t, ReadJSON(filepath, d))
	assert.Equal(t, "abc", d.Name)

	assert.NotNil(t, WriteJSON(filepath, make(chan int)))
	assert.NotNil(t, ReadJSON("./abc", d))
}

func TestMd5(t *testing.T) {
	res, err := Md5("../logo.png")
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	_, err2 := Md5("logo.png")
	assert.NotNil(t, err2)
}

func TestShadow(t *testing.T) {
	src := Temp("hello")
	assert.Equal(t, src+"_1", Shadow(src))
}
