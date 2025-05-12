package xapi

import (
	"net/http"

	"github.com/evercyan/brick/xhttp"
)

// ...
var (
	// Header 请求头
	Header = http.Header{
		"User-Agent": []string{xhttp.GetUserAgent()},
	}
)
