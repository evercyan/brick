package xtype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2List(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "a"}, String2List("a,b,a", ","))
	assert.Equal(t, []string{"a", "b"}, String2List("a,b,a", ",", true))
}
