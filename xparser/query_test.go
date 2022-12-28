package xparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURL(t *testing.T) {
	s := "http://baidu.com/query?a=1&b=2#c=3"
	u, err := ParseURL(s)
	assert.Nil(t, err)
	assert.Equal(t, "http", u.Scheme)
	assert.Equal(t, "baidu.com", u.Host)
	assert.Equal(t, "/query", u.Path)
	assert.Equal(t, "a=1&b=2", u.RawQuery)
	assert.Equal(t, "c=3", u.Fragment)
}

func TestParseQuery(t *testing.T) {
	s := "a=1&b=2&c="
	q := ParseQuery(s)
	assert.Equal(t, 3, len(q))
	assert.Equal(t, "1", q["a"])
	assert.Equal(t, map[string]string{}, ParseQuery("%"))
}

func TestBuildQuery(t *testing.T) {
	m := map[string]interface{}{
		"c": 2,
		"a": "a",
		"b": 1,
	}
	assert.Equal(t, "a=a&b=1&c=2", BuildQuery(m))
}
