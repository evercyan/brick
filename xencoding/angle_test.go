package xencoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrd(t *testing.T) {
	assert.Equal(t, 97, Ord("a"))
	assert.Equal(t, "a", Chr(97))
}

func TestHalfAngle(t *testing.T) {
	assert.Equal(t, ".", HalfAngle("。"))
}

func TestFullAngle(t *testing.T) {
	assert.Equal(t, "。", FullAngle("."))
}
