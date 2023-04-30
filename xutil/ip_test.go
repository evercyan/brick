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
}
