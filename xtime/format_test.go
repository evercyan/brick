package xtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tt, err := Parse("2024-01-01 01:02:03")
	assert.Nil(t, err)
	assert.Equal(t, 2024, tt.Year())

	assert.Equal(t, "2024-01-01 01:02:03", Format(tt, DateTime))

	assert.Equal(t, "2024-01-01 01:02:03", Format(tt, "his"))
}
