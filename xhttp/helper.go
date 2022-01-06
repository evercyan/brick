package xhttp

import (
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/evercyan/brick/xutil"
)

// getDecompressBody
func getDecompressBody(compressType string, body io.ReadCloser) (io.ReadCloser, error) {
	switch compressType {
	case "gzip":
		reader, err := gzip.NewReader(body)
		if err != nil {
			return nil, err
		}
		return reader, nil
	default:
		return body, nil
	}
}

// getRequestBody ...
func getRequestBody(body interface{}) []byte {
	if s, ok := body.(string); ok {
		return []byte(s)
	} else if b, ok := body.([]byte); ok {
		return b
	} else if xutil.IsJSONObject(body) {
		b, err := json.Marshal(body)
		if err != nil {
			return nil
		}
		return b
	}
	return nil
}
