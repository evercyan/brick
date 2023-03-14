package xregex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	assert.True(t, HasIP("abc127.0.0.1"))
	assert.False(t, HasIP("abc127.0"))
	assert.True(t, HasPhone("--18500000000"))
	assert.True(t, HasEmail("--e@q.com"))
	assert.True(t, HasLink("--http://b.com"))
	assert.True(t, HasDate("--2006-01-02"))
	assert.True(t, HasTime("--15:04:05"))
	assert.True(t, HasChinese("123哈哈456"))
}
