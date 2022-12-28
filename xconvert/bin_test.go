package xconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBin(t *testing.T) {
	assert.Equal(t, "00000000000000000000000000000001", Bin(1))
}
