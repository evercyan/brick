package xhttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	var (
		resStr    = `{"code":0}`
		doSuccess = func() (*http.Response, error) {
			return &http.Response{
				Body:   io.NopCloser(bytes.NewBuffer([]byte(resStr))),
				Header: nil,
			}, nil
		}
		doError = func() (*http.Response, error) {
			return nil, errors.New("error")
		}
	)

	assert.Equal(t, []byte(resStr), ToBytes(doSuccess()))
	assert.Equal(t, []byte(nil), ToBytes(doError()))

	assert.Equal(t, resStr, ToString(doSuccess()))
	assert.Equal(t, "", ToString(doError()))
}
