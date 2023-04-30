package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "96e79218965eb72c92a549dd5a330112", Md5("111111"))
}

func TestHash(t *testing.T) {
	assert.Less(t, uint64(0), Hash("abc", 1))
	assert.NotEmpty(t, HmacSha256("abc", "def"))
	assert.Equal(t, "7110eda4d09e062aa5e4a390b0a572ac0d2c0220", Sha1("1234"))
	assert.Equal(t, "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4", Sha256("1234"))
}
