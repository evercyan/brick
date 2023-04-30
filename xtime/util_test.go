package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tt, err := Parse("2021-01-01", "y-m-d")
	assert.Nil(t, err)
	assert.Equal(t, 2021, tt.Year())

	assert.Equal(t, "2021-01-01 00:00:00", Format(tt, "y-m-d h:i:s"))

	// 2021-01-01 00:00:00
	assert.Equal(t, int64(1609430400), First(tt))
	// 2021-01-01 23:59:59
	assert.Equal(t, int64(1609516799), Last(tt))

	assert.True(t, IsLeapYear(2000))

	var t1 time.Time
	assert.Less(t, int64(0), First(t1))
}

func TestCheck(t *testing.T) {
	assert.True(t, Check(2022, 1, 1))
	assert.True(t, Check(2022, 4, 30))

	assert.True(t, Check(2022, 2, 28))
	assert.False(t, Check(2022, 2, 29))

	assert.True(t, Check(2000, 2, 29))
	assert.False(t, Check(2000, 2, 30))

	assert.False(t, Check(2022, 4, 31))
	assert.False(t, Check(2022, 1, 32))
}
