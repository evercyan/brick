package xutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	obj := map[string]interface{}{
		"name": "hello",
		"list": []interface{}{
			map[string]interface{}{
				"age": 10,
			},
		},
		"nums": []int{
			1, 2, 3, 4,
		},
		"struct": struct{}{},
	}
	assert.Equal(t, "hello", Parse(obj, "name"))
	assert.Equal(t, 10, Parse(obj, "list.0.age"))
	assert.Equal(t, 2, Parse(obj, "nums.1"))
	assert.Equal(t, obj, Parse(obj, ""))
	assert.Equal(t, nil, Parse(obj, "list.0.aaa"))
	assert.Equal(t, nil, Parse(obj, "nums.abc"))
	assert.Equal(t, nil, Parse(obj, "nums.10"))
	assert.Equal(t, nil, Parse(obj, "struct.abc"))
}
