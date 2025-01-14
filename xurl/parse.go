package xurl

import (
	"fmt"
	"net/url"
)

// Parse ...
func Parse(str string) (*url.URL, error) {
	return url.Parse(str)
}

// Path ..
func Path(str string) string {
	u, err := Parse(str)
	if err != nil {
		return ""
	}
	return u.Path
}

// Domain ..
func Domain(str string) string {
	u, err := Parse(str)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

// Query ...
func Query(str string) map[string]string {
	res := make(map[string]string)
	values, err := url.ParseQuery(str)
	if err != nil {
		return nil
	}
	for k, v := range values {
		res[k] = v[0]
	}
	return res
}
