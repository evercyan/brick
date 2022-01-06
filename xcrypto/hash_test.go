package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Less(t, uint64(0), Hash("abc", 1))
	assert.NotEmpty(t, HmacSha256("abc", "def"))
}
