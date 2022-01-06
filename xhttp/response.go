package xhttp

import (
	"io"
	"net/http"
)

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
