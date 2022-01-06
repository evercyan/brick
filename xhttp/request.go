package xhttp

import (
	"context"
	"net/http"
)

// Get ...
func Get(ctx context.Context, url string) string {
	return ToString(NewClient().Get(ctx, url, nil))
}

// Post ...
func Post(ctx context.Context, url string, body interface{}) string {
	return ToString(NewClient().Post(ctx, url, http.Header{
		"Content-Type": []string{
			"application/json",
		},
	}, getRequestBody(body)))
}
