package xerror

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.Equal(t, 0, len(Errors()))
	err1 := New(1, "err1")
	New(2, "err2")
	New(3, "err3")
	assert.Equal(t, 3, len(Errors()))
	assert.Equal(t, err1, Errors()[0])
}

func TestNew(t *testing.T) {
	err1 := New(1, "err1")
	assert.Equal(t, 1, err1.Code)
	assert.Equal(t, "err1", err1.Msg)

	err2 := err1.WithMsg("err2")
	assert.Equal(t, 1, err2.Code)
	assert.Equal(t, `{"code":1,"msg":"err2"}`, err2.Error())
}
