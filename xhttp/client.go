package xhttp

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/http"
	"strings"
	"sync"
)

// Client ...
type Client struct {
	*http.Client

	tlsConfig *tls.Config
	option    *Option
}

// Get ...
func (t *Client) Get(
	ctx context.Context,
	url string,
	headers http.Header,
) (*http.Response, error) {
	return t.Do(ctx, "GET", url, headers, nil)
}

// Post ...
func (t *Client) Post(
	ctx context.Context,
	url string,
	headers http.Header,
	body []byte,
) (*http.Response, error) {
	return t.Do(ctx, "POST", url, headers, body)
}

// Put ...
func (t *Client) Put(
	ctx context.Context,
	url string,
	headers http.Header,
	body []byte,
) (*http.Response, error) {
	return t.Do(ctx, "PUT", url, headers, body)
}

// Delete ...
func (t *Client) Delete(
	ctx context.Context,
	url string,
	headers http.Header,
) (*http.Response, error) {
	return t.Do(ctx, "DELETE", url, headers, nil)
}

// Do ...
func (t *Client) Do(
	ctx context.Context,
	method string,
	url string,
	headers http.Header,
	body []byte,
) (*http.Response, error) {
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
	respBody, err := GetDecompressBody(res.Header.Get("Content-Encoding"), res.Body)
	if err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body = respBody

	return res, nil
}

// ----------------------------------------------------------------

var (
	defaultClient     *Client
	defaultClientOnce sync.Once
)

// New ...
func New() *Client {
	defaultClientOnce.Do(func() {
		defaultClient = NewWithtions(defaultOption)
	})
	return defaultClient
}

// NewWithtions ...
func NewWithtions(option *Option) *Client {
	option = setOptionDefaultValue(option)
	client := &Client{
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
