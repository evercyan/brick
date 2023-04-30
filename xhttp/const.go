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
)
