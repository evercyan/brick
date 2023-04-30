package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/evercyan/brick/xfile"
)

// ParseURL ...
func ParseURL(str string) (*url.URL, error) {
	return url.Parse(str)
}

// ParseQuery ...
func ParseQuery(str string) map[string]string {
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

// BuildFormData ...
func BuildFormData(header http.Header, m map[string]interface{}) (http.Header, io.Reader) {
	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	for k, v := range m {
		if vv, ok := v.(string); ok &&
			strings.HasPrefix(vv, "@") &&
			xfile.IsExist(strings.TrimPrefix(vv, "@")) {
			fpath := strings.TrimPrefix(vv, "@")
			f, err := os.Open(fpath)
			if err != nil {
				continue
			}
			part, err := w.CreateFormFile(k, filepath.Base(fpath))
			if err != nil {
				f.Close()
				continue
			}
			io.Copy(part, f)
			f.Close()
			continue
		}
		w.WriteField(k, fmt.Sprintf("%v", v))
	}
	defer w.Close()
	header.Set(HeaderKeyContentType, w.FormDataContentType())
	return header, b
}

// BuildReader ...
func BuildReader(data interface{}, types ...string) io.Reader {
	if r, ok := data.(io.Reader); ok {
		return r
	}
	if s, ok := data.(string); ok {
		return strings.NewReader(s)
	}
	if b, ok := data.([]byte); ok {
		return bytes.NewBuffer(b)
	}
	contentType := ""
	if len(types) > 0 {
		contentType = types[0]
	}
	switch contentType {
	case HeaderKeyContentTypeValueForm:
		if v, ok := data.(url.Values); ok {
			return strings.NewReader(v.Encode())
		}
		if m, ok := data.(map[string]interface{}); ok {
			return strings.NewReader(BuildValues(m).Encode())
		}
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil
		}
		return bytes.NewBuffer(b)
	}
	return nil
}

// Get ...
func Get(ctx context.Context, url string) string {
	res, err := New().Get(ctx, url, nil)
	if err != nil {
		return ""
	}
	return res.String()
}

// Post ...
func Post(ctx context.Context, url string, data interface{}) string {
	res, err := New().Post(ctx, url, nil, data)
	if err != nil {
		return ""
	}
	return res.String()
}

// ToString ...
func ToString(r *Response, err error) string {
	if err != nil {
		return ""
	}
	return r.String()
}

// ToResponse ...
func ToResponse(r *Response, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}
	return r.Response, nil
}
