package xhttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildReader(t *testing.T) {
	assert.NotNil(t, BuildReader("1"))
}
