package xtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDate(t *testing.T) {
	assert.True(t, CheckDate(2022, 1, 1))
	assert.True(t, CheckDate(2022, 4, 30))

	assert.True(t, CheckDate(2022, 2, 28))
	assert.False(t, CheckDate(2022, 2, 29))

	assert.True(t, CheckDate(2000, 2, 29))
	assert.False(t, CheckDate(2000, 2, 30))

	assert.False(t, CheckDate(2022, 4, 31))
	assert.False(t, CheckDate(2022, 1, 32))
}
