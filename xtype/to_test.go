package xtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo(t *testing.T) {
	assert.Equal(t, int64(0), ToInt64(nil))
	assert.Equal(t, int64(1), ToInt64(int(1)))
	assert.Equal(t, int64(1), ToInt64(int8(1)))
	assert.Equal(t, int64(1), ToInt64(int16(1)))
	assert.Equal(t, int64(1), ToInt64(int32(1)))
	assert.Equal(t, int64(1), ToInt64(int64(1)))
	assert.Equal(t, int64(1), ToInt64(uint(1)))
	assert.Equal(t, int64(1), ToInt64(uint8(1)))
	assert.Equal(t, int64(1), ToInt64(uint16(1)))
	assert.Equal(t, int64(1), ToInt64(uint32(1)))
	assert.Equal(t, int64(1), ToInt64(uint64(1)))
	assert.Equal(t, int64(1), ToInt64(float32(1)))
	assert.Equal(t, int64(1), ToInt64(float64(1)))
	assert.Equal(t, int64(1), ToInt64(true))
	assert.Equal(t, int64(0), ToInt64(false))
	assert.Equal(t, int64(0), ToInt64(""))
	assert.Equal(t, int64(0), ToInt64("abc"))
	assert.Equal(t, int64(1), ToInt64("1"))
	assert.Equal(t, int64(0), ToInt64(make([]int, 0)))
	assert.Equal(t, int(0), ToInt(0))
	assert.Equal(t, uint(0), ToUint(0))
	assert.Equal(t, uint64(0), ToUint64(0))
}

func TestToFloat64(t *testing.T) {
	assert.Equal(t, float64(0), ToFloat64(nil))
	assert.Equal(t, float64(1), ToFloat64(int(1)))
	assert.Equal(t, float64(1), ToFloat64(int8(1)))
	assert.Equal(t, float64(1), ToFloat64(int16(1)))
	assert.Equal(t, float64(1), ToFloat64(int32(1)))
	assert.Equal(t, float64(1), ToFloat64(int64(1)))
	assert.Equal(t, float64(1), ToFloat64(uint(1)))
	assert.Equal(t, float64(1), ToFloat64(uint8(1)))
	assert.Equal(t, float64(1), ToFloat64(uint16(1)))
	assert.Equal(t, float64(1), ToFloat64(uint32(1)))
	assert.Equal(t, float64(1), ToFloat64(uint64(1)))
	assert.Equal(t, float64(1), ToFloat64(float32(1)))
	assert.Equal(t, float64(1), ToFloat64(float64(1)))
	assert.Equal(t, float64(1), ToFloat64(true))
	assert.Equal(t, float64(0), ToFloat64(false))
	assert.Equal(t, float64(0), ToFloat64(""))
	assert.Equal(t, float64(0), ToFloat64("abc"))
	assert.Equal(t, float64(1), ToFloat64("1"))
	assert.Equal(t, float64(0), ToFloat64(make([]int, 0)))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "1234", ToString(1234))
	assert.Equal(t, "[1,2,3,4]", ToString([]int{1, 2, 3, 4}))
	assert.Equal(t, "haha", ToString("haha"))
	assert.Equal(t, "", ToString(nil))
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool(true))
	assert.False(t, ToBool("abc"))
	assert.False(t, ToBool("0"))
	assert.True(t, ToBool("1"))
	assert.False(t, ToBool(nil))
}

func TestToSlice(t *testing.T) {
	assert.Equal(t, []interface{}{}, ToSlice(nil))
	assert.Equal(t, []interface{}{1}, ToSlice(1))
	assert.Equal(t, []interface{}{"a"}, ToSlice("a"))
	assert.Equal(t, []interface{}{"a", "b", "c"}, ToSlice("a, b, , ,c"))
	assert.Equal(t, []interface{}{"a", "b"}, ToSlice([]interface{}{"a", "b"}))
	assert.Equal(t, []interface{}{1, 2, 3}, ToSlice([]int{1, 2, 3}))
	assert.Equal(t, []interface{}{1, 2, 3}, ToSlice([3]int{1, 2, 3}))
	assert.Equal(t, []interface{}{1}, ToSlice(map[string]int{"a": 1}))
	assert.Equal(t, []interface{}{}, ToSlice(""))
}
