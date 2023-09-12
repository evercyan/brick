package xhttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/evercyan/brick/xjson"
	"github.com/stretchr/testify/assert"
)

var (
	prefix = "https://httpbin.org"
	ctx    = context.Background()
)

func TestRequest(t *testing.T) {
	client := New()

	_, err1 := client.Get(ctx, prefix+"/get", nil)
	assert.Nil(t, err1)

	_, err2 := client.Put(ctx, prefix+"/put", nil, nil)
	assert.Nil(t, err2)

	_, err3 := client.Delete(ctx, prefix+"/delete", nil)
	assert.Nil(t, err3)

	_, err4 := client.Patch(ctx, prefix+"/patch", nil)
	assert.Nil(t, err4)

	_, err5 := client.Get(ctx, prefix+"/gzip", nil)
	assert.Nil(t, err5)
}

func TestPost(t *testing.T) {
	client := New()

	// header 默认为 application/json
	// body 会通过 BuildReader 处理成对应数据
	// string, bytes, io.Reader, url.Values
	// json.Marshal
	body := map[string]interface{}{
		"name":   "brick",
		"number": 1,
	}

	// JSON 请求
	resp1, err1 := client.Post(ctx, prefix+"/post", nil, body)
	assert.Nil(t, err1)
	if resp1 != nil {
		fmt.Println(xjson.Pretty(resp1.String()))
	}
	/*
		{
		    "args": {},
		    "data": "{\"name\":\"brick\",\"number\":1}",
		    "files": {},
		    "form": {},
		    "headers": {
		        "Accept-Encoding": "gzip",
		        "Content-Length": "27",
		        "Content-Type": "application/json",
		        "Host": "httpbin.org",
		        "User-Agent": "Go-http-client/2.0",
		        "X-Amzn-Trace-ID": "Root=1-6444ff64-5d0f1a6e2a1ca2812bb8d9e2"
		    },
		    "json": {
		        "name": "brick",
		        "number": 1
		    },
		    "origin": "103.88.46.208",
		    "url": "https://httpbin.org/post"
		}
	*/

	// Form 表单提交
	header1 := http.Header{}
	header1.Set(HeaderKeyContentType, HeaderKeyContentTypeValueForm)
	resp2, err2 := client.Post(ctx, prefix+"/post", header1, body)
	assert.Nil(t, err2)
	if resp2 != nil {
		fmt.Println(xjson.Pretty(resp2.String()))
	}
	/*
		{
		    "args": {},
		    "data": "",
		    "files": {},
		    "form": {
		        "name": "brick",
		        "number": "1"
		    },
		    "headers": {
		        "Accept-Encoding": "gzip",
		        "Content-Length": "19",
		        "Content-Type": "application/x-www-form-urlencoded",
		        "Host": "httpbin.org",
		        "User-Agent": "Go-http-client/2.0",
		        "X-Amzn-Trace-ID": "Root=1-6444ff65-6fbdbb991b5eef675e2c4c82"
		    },
		    "json": null,
		    "origin": "103.88.46.208",
		    "url": "https://httpbin.org/post"
		}
	*/
}

func TestUploadFile(t *testing.T) {
	header := http.Header{}
	header.Set(HeaderKeyContentType, HeaderKeyContentTypeValueFormData)
	resp3, err3 := New().Post(ctx, prefix+"/post", header, map[string]interface{}{
		"name": "brick",
		"file": "@./README.md",
	})
	assert.Nil(t, err3)
	if resp3 != nil {
		fmt.Println(xjson.Pretty(resp3.String()))
	}
	/*
		{
		    "args": {},
		    "data": "",
		    "files": {
		        "file": "# xhttp\n\n- [ ] \u652f\u6301\u5bfc\u51facurl\u547d\u4ee4\n\n"
		    },
		    "form": {
		        "name": "brick"
		    },
		    "headers": {
		        "Accept-Encoding": "gzip",
		        "Content-Length": "398",
		        "Content-Type": "multipart/form-data; boundary=8ddbc5f48344bbfe44818557a3d7b567163a7493b38a297de14447efc788",
		        "Host": "httpbin.org",
		        "User-Agent": "Go-http-client/2.0",
		        "X-Amzn-Trace-ID": "Root=1-6444ff3b-1912fb5716c711334c613b09"
		    },
		    "json": null,
		    "origin": "103.88.46.208",
		    "url": "https://httpbin.org/post"
		}
	*/
}

func TestTrace(t *testing.T) {
	client := New(WithTraceEnable())
	resp, err := client.Get(ctx, prefix+"/get", nil)
	assert.Nil(t, err)

	trace := resp.Trace()
	assert.NotNil(t, trace)
	if trace != nil {
		fmt.Println(xjson.Pretty(trace))
		fmt.Println(trace.TotalTime.String())
	}
}

func TestCookies(t *testing.T) {
	header := http.Header{}
	header.Set("Cookie", "name=abc")
	resp, err := New().Get(ctx, prefix+"/cookies/set/hello/world", header)
	assert.Nil(t, err)
	if resp != nil {
		fmt.Println(xjson.Pretty(resp.Header))
		fmt.Println(xjson.Pretty(resp.String()))
	}
	/*
		{
		    "gzipped": true,
		    "headers": {
		        "Accept-Encoding": "gzip",
		        "Content-Length": "4",
		        "Content-Type": "application/json",
		        "Host": "httpbin.org",
		        "User-Agent": "Go-http-client/2.0",
		        "X-Amzn-Trace-ID": "Root=1-644deada-3512ae9802549aa317f5916c"
		    },
		    "method": "GET",
		    "origin": "45.90.208.27"
		}
	*/
}

func TestGzip(t *testing.T) {
	resp, err := New().Get(ctx, prefix+"/gzip", nil)
	assert.Nil(t, err)
	if resp != nil {
		fmt.Println(xjson.Pretty(resp.String()))
	}
}

func TestCoverage(t *testing.T) {
	client := New(
		WithTraceEnable(),
		WithRetryTimes(3),
		WithRequestTimeout(20*time.Second),
		WithDialer(&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}),
		WithTransport(&http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			TLSHandshakeTimeout: 10 * time.Second,
		}),
	)
	res, err := client.Get(ctx, prefix+"/get", nil)
	assert.Nil(t, err)

	assert.NotEmpty(t, Get(ctx, prefix+"/get"))
	assert.NotEmpty(t, Post(ctx, prefix+"/post", nil))
	assert.NotEmpty(t, ToString(res, err))
	resp, _ := ToResponse(res, err)
	assert.NotEmpty(t, resp)
}
