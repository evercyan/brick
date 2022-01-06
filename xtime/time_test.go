package xtime

import (
	"testing"

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
}
