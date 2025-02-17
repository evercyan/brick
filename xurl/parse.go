package xurl

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xlodash"
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

// Name ...
func Name(url string, exts ...bool) string {
	upath := Path(url)
	if upath == "" {
		return ""
	}
	uname := upath
	parts := strings.Split(upath, "/")
	if len(parts) > 1 {
		uname = parts[len(parts)-1]
	}
	if !xlodash.First(exts, true) {
		uname = xfile.GetFileNameWithoutExt(uname)
	}
	return uname
}

// FullName ...
func FullName(url string, exts ...bool) string {
	upath := Path(url)
	if upath == "" {
		return ""
	}
	uname := strings.Trim(strings.ReplaceAll(upath, "/", "_"), "_")
	if !xlodash.First(exts, true) {
		uname = xfile.GetFileNameWithoutExt(uname)
	}
	return uname
}
