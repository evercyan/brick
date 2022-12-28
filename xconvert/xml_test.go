package xconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXML2Map(t *testing.T) {
	str := `<xml>
				<a>a1</a>
				<b>b1</b>
				<c><c1>c11</c1></c>
			</xml>`

	m, err := XML2Object(str)

	assert.Nil(t, err)
	assert.Equal(t, "a1", m["a"])
	assert.Equal(t, "b1", m["b"])
	assert.Equal(t, "", m["c"])
}
