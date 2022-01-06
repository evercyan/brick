package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "96e79218965eb72c92a549dd5a330112", Md5("111111"))
}
