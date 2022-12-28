package xregex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	assert.True(t, IsIPV4("127.0.0.1"))
	assert.True(t, IsIPV4("255.255.255.255"))
	assert.False(t, IsIPV4("256.255.255.255"))
	assert.False(t, IsIPV6("127.0.0.1"))
	assert.True(t, IsIPV6("2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b"))
	assert.True(t, IsIP("127.0.0.1"))

	assert.True(t, IsMacAddress("6a:00:02:9c:1f:30"))
	assert.False(t, IsMacAddress(""))

	assert.True(t, IsPhone("18500000000"))
	assert.False(t, IsPhone("134"))

	assert.True(t, IsEmail("e@q.com"))
	assert.False(t, IsEmail("eq.com"))

	assert.True(t, IsLink("http://b.com"))
	assert.True(t, IsLink("b.com"))
	assert.False(t, IsLink("com"))

	assert.True(t, IsDate("2006-01-02"))
	assert.True(t, IsTime("15:04:05"))
	assert.True(t, IsHexColor("#000000"))

	assert.True(t, IsIdcard("342622200001010101"))
	assert.True(t, IsIdcard("34262220000101010x"))
	assert.True(t, IsIdcard("34262220000101010X"))
	assert.False(t, IsIdcard("3426222000010101"))
}
