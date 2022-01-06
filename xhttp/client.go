package xhttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"sync"
)

// client ...
type client struct {
	*http.Client
	tlsConfig *tls.Config
	option    *Option
}

// Get ...
func (t *client) Get(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	return t.Do(ctx, "GET", url, headers, nil)
}

// Post ...
func (t *client) Post(ctx context.Context, url string, headers http.Header, body []byte) (*http.Response, error) {
	return t.Do(ctx, "POST", url, headers, body)
}

// Put ...
func (t *client) Put(ctx context.Context, url string, headers http.Header, body []byte) (*http.Response, error) {
	return t.Do(ctx, "PUT", url, headers, body)
}

// Delete ...
func (t *client) Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	return t.Do(ctx, "DELETE", url, headers, nil)
}

// Do ...
func (t *client) Do(ctx context.Context, method string, url string, headers http.Header, body []byte) (*http.Response, error) {
	// https
	if strings.HasPrefix(url, "https") {
		if transport, ok := t.Client.Transport.(*http.Transport); ok {
			transport.TLSClientConfig = t.tlsConfig
		}
	}

	// header
	if headers == nil {
		headers = make(http.Header)
	}
	if _, ok := headers["Accept"]; !ok {
		headers["Accept"] = []string{"*/*"}
	}
	if _, ok := headers["Accept-Encoding"]; !ok && t.option.Compressed {
		headers["Accept-Encoding"] = []string{"deflate, gzip"}
	}

	// request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = headers

	// inject trace

	// retry request
	var (
		res    *http.Response
		resErr error
	)
	for i := 0; i < t.option.RetryTimes; i++ {
		res, resErr = t.Client.Do(req)
		if resErr == nil {
			break
		}
	}
	if resErr != nil {
		return nil, resErr
	}

	// compress
	respBody, err := getDecompressBody(res.Header.Get("Content-Encoding"), res.Body)
	if err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body = respBody

	return res, nil
}

// ----------------------------------------------------------------

var (
	defaultClient     *client
	defaultClientOnce sync.Once
)

// NewClient ...
func NewClient() *client {
	defaultClientOnce.Do(func() {
		defaultClient = NewClientWithOptions(defaultOption)
	})
	return defaultClient
}

// NewClientWithOptions ...
func NewClientWithOptions(option *Option) *client {
	option = setOptionDefaultValue(option)
	client := &client{
		Client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   option.ConnsPerHost,
				TLSHandshakeTimeout:   option.HandshakeTimeout,
				ResponseHeaderTimeout: option.ResponseHeaderTimeout,
				DisableCompression:    !option.Compressed,
			},
		},
		option: option,
	}
	if option.SSLEnabled {
		client.Client.Timeout = option.RequestTimeout
		client.tlsConfig = option.TLSConfig
	}
	return client
}
