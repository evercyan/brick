package xgen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	assert.NotEmpty(t, GUID())
	fmt.Println("GUID:", GUID())

	assert.NotEmpty(t, UUID())
	fmt.Println("UUID:", UUID())

	assert.NotEmpty(t, XID())
	fmt.Println("XID:", XID())

	assert.NotEmpty(t, Nanoid())
	fmt.Println("Nanoid:", Nanoid())
}
