package xcolor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOption(t *testing.T) {
	for s := StyNone; s <= StyCrossedOut; s++ {
		assert.NotEmpty(t, s.String())
	}
	for fc := FgNone; fc <= FgWhite; fc++ {
		assert.NotEmpty(t, fc.String())
	}
	for bc := BgNone; bc <= BgWhite; bc++ {
		assert.NotEmpty(t, bc.String())
	}
}
