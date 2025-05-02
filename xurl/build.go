package xurl

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// BuildURL ...
func BuildURL(url string, m map[string]interface{}) string {
	list := make([]string, 0)
	for k, v := range m {
		list = append(list, fmt.Sprintf("%s=%v", k, v))
	}
	sort.Strings(list)
	query := strings.Join(list, "&")
	if url == "" {
		return query
	}
	symbol := "?"
	if strings.Contains(url, "?") {
		symbol = "&"
	}
	return url + symbol + query
}

// BuildValues ...
func BuildValues(m map[string]interface{}) url.Values {
	res := make(url.Values)
	for k, v := range m {
		res[k] = []string{fmt.Sprint(v)}
	}
	return res
}

// BuildCookie ...
func BuildCookie(cookies map[string]string) string {
	pairs := make([]string, 0)
	for k, v := range cookies {
		pairs = append(pairs, k+"="+v)
	}
	return strings.Join(pairs, ";")
}
