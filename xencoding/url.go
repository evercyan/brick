package xencoding

import (
	"net/url"
)

// URLEncode ...
func URLEncode(text string) string {
	return url.QueryEscape(text)
}

// URLDecode ...
func URLDecode(text string) string {
	resp, err := url.QueryUnescape(text)
	if err != nil {
		return ""
	}
	return resp
}
