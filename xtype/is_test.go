package xtype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.True(t, IsInt(1))
	assert.True(t, IsUint(uint32(1)))
	assert.True(t, IsFloat(float32(1)))

	assert.True(t, IsNumeric(1))
	assert.True(t, IsNumeric(uint32(1)))
	assert.True(t, IsNumeric(float32(1)))

	assert.True(t, IsBool(true))
	assert.False(t, IsBool("1"))

	assert.True(t, IsString("hello"))
	assert.False(t, IsString(1))

	assert.True(t, IsSlice([]int{1, 2, 3}))
	assert.True(t, IsArray([3]int{1, 2, 3}))
	assert.True(t, IsMap(make(map[string]string)))
	assert.True(t, IsChannel(make(chan string)))

	assert.True(t, IsTime(time.Now()))
	assert.False(t, IsTime("1"))

	elemStruct := struct {
		Name string
	}{}
	assert.True(t, IsStruct(elemStruct))

	elemFunc := func() {}
	assert.True(t, IsFunc(elemFunc))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.True(t, IsEmpty(0))
	assert.True(t, IsEmpty(nil))
	assert.True(t, IsEmpty(false))
	assert.False(t, IsEmpty("0"))
}

func TestIsJson(t *testing.T) {
	assert.True(t, IsJSONString("{\"Title\":\"AAA\",\"Text\":\"BBB\"}"))
	assert.True(t, IsJSONString("[1, 2, 3]"))
	assert.False(t, IsJSONString("["))
	assert.True(t, IsJSONObject(struct{}{}))
	assert.False(t, IsJSONObject(1))
}

func TestIsContains(t *testing.T) {
	assert.True(t, IsContains([]int{1, 2, 3}, 1))
	assert.False(t, IsContains([]int{1, 2, 3}, 4))
	assert.True(t, IsContains(map[int]int{1: 1, 2: 2}, 1))
	assert.False(t, IsContains(map[int]int{1: 1, 2: 2}, 3))
}
