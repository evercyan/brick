package xhttp

import (
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/evercyan/brick/xtype"
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
	if s, sok := body.(string); sok {
		return []byte(s)
	} else if b, bok := body.([]byte); bok {
		return b
	} else if xtype.IsJSONObject(body) {
		b1, err := json.Marshal(body)
		if err != nil {
			return nil
		}
		return b1
	}
	return nil
}
