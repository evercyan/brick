package xhttp

// gomonkey
// MAC OS X arm64
// https://github.com/agiledragon/gomonkey/issues/19

// import (
// 	"bytes"
// 	"context"
// 	"io"
// 	"net/http"
// 	"reflect"
// 	"testing"
//
// 	"github.com/agiledragon/gomonkey/v2"
// 	"github.com/stretchr/testify/assert"
// )
//
// // gomonkey mock
// // go test -gcflags 'all=-N -l' ./
//
// var (
// 	ctx    = context.Background()
// 	reqUrl = "http://hello.world"
// 	resStr = `{"code":0}`
// )
//
// // TestNewClient ...
// func TestNewClient(t *testing.T) {
//
// 	// mock (*http.Client).Do
// 	var c *http.Client
// 	patches := gomonkey.ApplyMethod(reflect.TypeOf(c), "Do", func(_ *http.Client, req *http.Request) (*http.Response, error) {
// 		return &http.Response{
// 			Body:   io.NopCloser(bytes.NewBuffer([]byte(resStr))),
// 			Header: nil,
// 		}, nil
// 	})
// 	defer patches.Reset()
//
// 	// client
// 	client := NewClient()
//
// 	// Get
// 	res, err := client.Get(ctx, reqUrl, nil)
// 	assert.Nil(t, err)
// 	if err == nil {
// 		defer res.Body.Close()
// 		b, _ := io.ReadAll(res.Body)
// 		assert.Equal(t, resStr, string(b))
// 	}
//
// }
