package xurl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// e.g. http://baidu.com/query?a=1&b=2#c=3
	// Scheme 		http
	// Host 		baidu.com
	// Path 		/query
	// RawQuery 	a=1&b=2
	// Fragment 	c=3
	s := "http://baidu.com/path?a=1&b=2#c=3"
	u, err := Parse(s)
	assert.Nil(t, err)
	assert.Equal(t, "http", u.Scheme)
	assert.Equal(t, "baidu.com", u.Host)
	assert.Equal(t, "/query", u.Path)
	assert.Equal(t, "a=1&b=2", u.RawQuery)
	assert.Equal(t, "c=3", u.Fragment)

	query := Query(u.RawQuery)
	assert.Equal(t, 2, len(query))
	assert.Equal(t, "1", query["a"])

	assert.Equal(t, "path", Path(s))
	assert.Equal(t, "http://baidu.com", Domain(s))
}

func TestName(t *testing.T) {
	url := "http://baidu.com/a.a/b.b/c.c.html"
	assert.Equal(t, "c.c.html", Name(url))
	assert.Equal(t, "c.c", Name(url, false))

	assert.Equal(t, "a.a_b.b_c.c.html", FullName(url))
	assert.Equal(t, "a.a_b.b_c.c", FullName(url, false))
}
