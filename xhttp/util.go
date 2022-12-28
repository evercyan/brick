package xhttp

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/evercyan/brick/xtype"
)

// Get ...
func Get(ctx context.Context, url string) string {
	return ToString(New().Get(ctx, url, nil))
}

// Post ...
func Post(ctx context.Context, url string, body interface{}) string {
	return ToString(New().Post(ctx, url, http.Header{
		"Content-Type": []string{
			"application/json",
		},
	}, Encode(body)))
}

// Encode ...
func Encode(body interface{}) []byte {
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

// ToBytes ...
func ToBytes(res *http.Response, err error) []byte {
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	return b
}

// ToString ...
func ToString(res *http.Response, err error) string {
	return string(ToBytes(res, err))
}

// GetDecompressBody ...
func GetDecompressBody(compressType string, body io.ReadCloser) (io.ReadCloser, error) {
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
