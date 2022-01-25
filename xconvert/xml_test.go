package xconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXML2Map(t *testing.T) {
	str := `<xml>
				<a>a</a>
				<c>c</c>
			</xml>`

	m, err := XML2Map(str)

	assert.Nil(t, err)
	assert.Equal(t, "a", m["a"])
	assert.Equal(t, "c", m["c"])
}
