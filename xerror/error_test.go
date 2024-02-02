package xerror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := New(1, "error")
	assert.Equal(t, 1, err.Code)
}

func TestIsError(t *testing.T) {
	assert.True(t, IsXErr(New(1, "error")))
	assert.False(t, IsXErr(errors.New("error")))
}

func TestStack(t *testing.T) {
	err := New(1, "error")
	fmt.Println(xjson.Encode(err))
}

func TestXErrWrap(t *testing.T) {
	err1 := errors.New("错误 1")
	err2 := New(2, "错误 2").Wrap(err1)
	err3 := New(3, "错误 3").Wrap(err2)

	fmt.Println("err1:", err1.Error())
	fmt.Println("err2:", err2.Error())
	fmt.Println("err3:", err3.Error())

	var err error = err3
	for {
		err = Unwrap(err)
		if err == nil {
			break
		}
		fmt.Println("Unwrap:", err.Error(), "xError:", IsXErr(err))
	}
}

func TestErrWrap(t *testing.T) {
	err1 := errors.New("error1")
	err2 := New(2, "error2")

	wrappedErr := errors.New("wrappedErr")

	werr1 := Wrap(err1, wrappedErr)
	fmt.Println("werr1", werr1.Error())
	fmt.Println("werr1 wrapped", Unwrap(werr1).Error())

	werr2 := Wrap(err2, wrappedErr)
	fmt.Println("werr2", werr2.Error())
	fmt.Println("werr2 wrapped", Unwrap(werr2).Error())
}
