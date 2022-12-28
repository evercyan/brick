package xjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var str = `{
    "name": "hello",
    "detail": {
    	"age": 20,
        "height": "175cm",
        "weight": "60kg"
    },
    "langs": [
        "php",
        "golang",
        "python",
        "shell"
    ]
}`

func TestJson(t *testing.T) {
	name := New(str).Key("name").ToString()
	assert.Equal(t, "hello", name)

	age := New(str).Key("detail").Key("age").ToInt()
	assert.Equal(t, int64(20), age)

	lang1 := New(str).Key("langs").Index(1).ToString()
	assert.Equal(t, "golang", lang1)

	langs1 := New(str).Key("langs").ToJSON()
	assert.Equal(t, `["php","golang","python","shell"]`, langs1)

	langs2 := New(str).Key("langs").ToArray()
	assert.ElementsMatch(t, []interface{}{"python", "shell", "php", "golang"}, langs2.([]interface{}))

	detail := New(str).Key("detail").ToArray()
	assert.Equal(t, map[string]interface{}{
		"age":    float64(20),
		"height": "175cm",
		"weight": "60kg",
	}, detail.(map[string]interface{}))

	assert.Nil(t, New(str).Key("name").ToArray())
}

func TestJsonError(t *testing.T) {
	assert.Empty(t, New("{").ToJSON())
	assert.Nil(t, New("{").Key("name").Value())
	assert.Empty(t, New("{").Key("name").ToString())
	assert.Equal(t, int64(0), New("{").Key("name").ToInt())
	assert.Nil(t, New(str).Key("name1").Value())
	assert.Nil(t, New(str).Key("langs").Key("name").Value())
	assert.Nil(t, New(str).Index(0).Value())
	assert.Nil(t, New(str).Key("langs").Index(10).Value())
	assert.Equal(t, int64(0), New(str).Key("langs").ToInt())
}

// BenchmarkJson-8   	  327823	      3381 ns/op	    1472 B/op	      37 allocs/op
func BenchmarkJSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		New(str).Key("name").ToString()
	}
}
