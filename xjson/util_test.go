package xjson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type js struct {
	Title string
	Text  string
}

func TestJSONEncode(t *testing.T) {
	j := &js{
		Title: "AAA",
		Text:  "BBB",
	}
	assert.JSONEq(t, `{"Text":"BBB","Title":"AAA"}`, Encode(j))
}

func TestJSONDecode(t *testing.T) {
	j := &js{}
	assert.Nil(t, Decode(`{"Text":"BBB","Title":"AAA"}`, j))
	assert.Equal(t, "AAA", j.Title)
}

func TestPretty(t *testing.T) {
	data := map[string]string{
		"hello": "world",
	}
	assert.NotEmpty(t, Pretty(data))
	assert.NotEmpty(t, Pretty(`{"hello": "world"}`))
	assert.NotEmpty(t, Pretty(`"{\"hello\": \"world\"}"`))
}

func TestMinify(t *testing.T) {
	j := &js{
		Title: "AAA",
		Text:  "BBB",
	}
	assert.JSONEq(t, `{"Text":"BBB","Title":"AAA"}`, Minify(j))

	s := `{
    "Text": "BBB",
    "Title": "AAA"
}`
	assert.JSONEq(t, `{"Text":"BBB","Title":"AAA"}`, Minify(s))
	assert.JSONEq(t, `{"hello":"world"}`, Minify(`"{\"hello\": \"world\"}"`))
}

func TestFormat(t *testing.T) {
	target := `{"name":"abc","list":["a", "b"]}`

	assert.Equal(t, target, Format(target))
	assert.Equal(t, target, Format(fmt.Sprintf("111%s222", target)))
	assert.Equal(t, target, Format(`{\"name\":\"abc\",\"list\":[\"a\", \"b\"]}`))
	assert.Equal(t, target, Format(`"{\"name\":\"abc\",\"list\":[\"a\", \"b\"]}"`))
}

func TestSort(t *testing.T) {
	assert.Equal(t, `[1,2,3,4,5]`, Sort(`[1,3,4,2,5]`))
	assert.Equal(t, `["1","2","3","4","5"]`, Sort(`["1","3","4","2","5"]`))
	assert.Equal(t, `["a","b","c","d"]`, Sort(`["c","a","d","b"]`))
	assert.Equal(t, `[1,"a","b","c","d"]`, Sort(`["c","a","d","b",1]`))
	assert.Equal(t, `{"name":"abc"}`, Sort(`{"name":"abc"}`))
	assert.Equal(t, `[{"name":"abc"}]`, Sort(`[{"name":"abc"}]`))
}

func TestFilter(t *testing.T) {
	assert.Equal(t, `{"name":"abc"}`, FilterUnquote(`{\"name\":\"abc\"}`))
	assert.Equal(t, `{"name":"abc"}`, FilterUnquote(`"{\"name\":\"abc\"}"`))
	assert.Equal(t, `{}`, FilterPrefix(`abc{}`))
	assert.Equal(t, `{}`, FilterSuffix(`{}abc`))
}
