package xhttp

import (
	"net/http"
)

// Method ...
const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

// HeaderKey header 字段, 同 http.CanonicalHeaderKey() 效果一致
const (
	HeaderKeyUserAgent       = "User-Agent"
	HeaderKeyAccept          = "Accept"
	HeaderKeyContentType     = "Content-Type"
	HeaderKeyContentLength   = "Content-Length"
	HeaderKeyContentEncoding = "Content-Encoding"
	HeaderKeyLocation        = "Location"
	HeaderKeyAuthorization   = "Authorization"
	HeaderKeyAcceptEncoding  = "Accept-Encoding"

	HeaderKeyContentTypeValueJSON     = "application/json"
	HeaderKeyContentTypeValueForm     = "application/x-www-form-urlencoded"
	HeaderKeyContentTypeValueFormData = "multipart/form-data"
	HeaderKeyContentEncodingValueGzip = "gzip"
)

// ...
var (
	HeaderContentTypeJSON = http.Header{
		HeaderKeyContentType: []string{
			HeaderKeyContentTypeValueJSON,
		},
	}
	HeaderContentTypeForm = http.Header{
		HeaderKeyContentType: []string{
			HeaderKeyContentTypeValueForm,
		},
	}
	HeaderContentTypeFormData = http.Header{
		HeaderKeyContentType: []string{
			HeaderKeyContentTypeValueFormData,
		},
	}
	UserAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
	}
)
