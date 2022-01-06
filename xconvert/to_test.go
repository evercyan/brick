package xconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	assert.Equal(t, uint64(0), ToUint(nil))
	assert.Equal(t, uint64(1), ToUint(int(1)))
	assert.Equal(t, uint64(1), ToUint(int8(1)))
	assert.Equal(t, uint64(1), ToUint(int16(1)))
	assert.Equal(t, uint64(1), ToUint(int32(1)))
	assert.Equal(t, uint64(1), ToUint(int64(1)))
	assert.Equal(t, uint64(1), ToUint(uint(1)))
	assert.Equal(t, uint64(1), ToUint(uint8(1)))
	assert.Equal(t, uint64(1), ToUint(uint16(1)))
	assert.Equal(t, uint64(1), ToUint(uint32(1)))
	assert.Equal(t, uint64(1), ToUint(uint64(1)))
	assert.Equal(t, uint64(1), ToUint(float32(1)))
	assert.Equal(t, uint64(1), ToUint(float64(1)))
	assert.Equal(t, uint64(1), ToUint(true))
	assert.Equal(t, uint64(0), ToUint(false))
	assert.Equal(t, uint64(0), ToUint(""))
	assert.Equal(t, uint64(0), ToUint("abc"))
	assert.Equal(t, uint64(1), ToUint("1"))
	assert.Equal(t, uint64(0), ToUint(make([]int, 0)))
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
	assert.Equal(t, "[1 2 3 4]", ToString([]int{1, 2, 3, 4}))
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool(true))
	assert.False(t, ToBool("abc"))
	assert.False(t, ToBool("0"))
	assert.True(t, ToBool("1"))
	assert.False(t, ToBool(nil))
}
