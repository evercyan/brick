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
}
