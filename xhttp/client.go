package xhttp

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

// Client ...
type Client struct {
	*http.Client

	option *Option
	trace  *Trace
}

// Get ...
func (t *Client) Get(ctx context.Context, url string, header http.Header) (*Response, error) {
	return t.Do(ctx, MethodGet, url, header, nil)
}

// Post ...
func (t *Client) Post(ctx context.Context, url string, header http.Header, data interface{}) (*Response, error) {
	return t.Do(ctx, MethodPost, url, header, data)
}

// Put ...
func (t *Client) Put(ctx context.Context, url string, header http.Header, data interface{}) (*Response, error) {
	return t.Do(ctx, MethodPut, url, header, data)
}

// Delete ...
func (t *Client) Delete(ctx context.Context, url string, header http.Header) (*Response, error) {
	return t.Do(ctx, MethodDelete, url, header, nil)
}

// Patch ...
func (t *Client) Patch(ctx context.Context, url string, header http.Header) (*Response, error) {
	return t.Do(ctx, MethodPatch, url, header, nil)
}

// Do ...
func (t *Client) Do(
	ctx context.Context, method string, url string, header http.Header, data interface{},
) (*Response, error) {
	defer func() {
		if t.trace != nil {
			t.trace.Finish()
		}
	}()
	if header == nil {
		header = http.Header{}
		header.Set(HeaderKeyContentType, HeaderKeyContentTypeValueJSON)
	}
	// trace
	if t.option.TraceEnable {
		t.trace = new(Trace)
		ctx = t.trace.WithClientTrace(ctx)
	}
	var (
		resp *http.Response
		err  error
	)
	var body io.Reader
	m, ok := data.(map[string]interface{})
	if ok && header.Get(HeaderKeyContentType) == HeaderKeyContentTypeValueFormData {
		header, body = BuildFormData(header, m)
	} else {
		body = BuildReader(data, header.Get(HeaderKeyContentType))
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = header
	// retry
	for i := 0; i < t.option.RetryTimes; i++ {
		resp, err = t.Client.Do(req)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	reader := resp.Body
	res := &Response{
		Response: resp,
		trace:    t.trace,
	}
	// gzip
	if resp.Header.Get(HeaderKeyContentEncoding) == HeaderKeyContentEncodingValueGzip {
		if _, ok := body.(*gzip.Reader); ok {
			reader, err = gzip.NewReader(reader)
			if err != nil {
				return nil, err
			}
			defer reader.Close()
		}
	}
	if res.body, err = io.ReadAll(reader); err != nil {
		return nil, err
	}
	return res, nil
}

// ----------------------------------------------------------------

// New ...
func New(options ...OptionFn) *Client {
	config := defaultOption
	for _, fn := range options {
		fn(config)
	}
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return &Client{
		Client: &http.Client{
			Timeout:   config.RequestTimeout,
			Transport: config.Transport,
			Jar:       cookieJar,
		},
		option: config,
	}
}
