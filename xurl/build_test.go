package xurl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildQuery(t *testing.T) {
	m := map[string]interface{}{
		"c": 2,
		"a": "a",
		"b": 1,
	}
	assert.Equal(t, "a=a&b=1&c=2", BuildURL("", m))
	assert.Equal(t, "http://abc?a=a&b=1&c=2", BuildURL("http://abc", m))
	assert.Equal(t, "http://abc?h=1&a=a&b=1&c=2", BuildURL("http://abc?h=1", m))
}

func TestBuildValues(t *testing.T) {
	m := map[string]interface{}{
		"c": 2,
		"a": "a",
		"b": 1,
	}
	assert.NotNil(t, BuildValues(m))
}
