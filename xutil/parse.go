package xutil

import (
	"net/url"
)

// ParseURL ...
func ParseURL(str string) (*url.URL, error) {
	// e.g. http://baidu.com/query?a=1&b=2#c=3
	// Scheme 		http
	// Host 		baidu.com
	// Path 		/query
	// RawQuery 	a=1&b=2
	// Fragment 	c=3
	return url.Parse(str)
}

// ParseQuery ...
func ParseQuery(str string) map[string]string {
	// e.g. a=1&b=2&c=3
	// map[string]string{
	// 		"a": "1",
	// 		"b": "2",
	// 		"c": "",
	// }
	res := make(map[string]string)
	values, err := url.ParseQuery(str)
	if err != nil {
		return res
	}
	for k, v := range values {
		res[k] = v[0]
	}
	return res
}
