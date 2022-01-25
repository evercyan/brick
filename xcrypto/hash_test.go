package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Less(t, uint64(0), Hash("abc", 1))
	assert.NotEmpty(t, HmacSha256("abc", "def"))
	assert.Equal(t, "7110eda4d09e062aa5e4a390b0a572ac0d2c0220", Sha1("1234"))
}
