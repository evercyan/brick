package xutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	fmt.Println("GetIntranetIP", GetIntranetIP())
	assert.NotNil(t, GetIntranetIP())

	fmt.Println("GetExternalIP", GetExternalIP())
	assert.NotNil(t, GetExternalIP())

	assert.False(t, IsExternalIP("127.0.0.1"))
}
