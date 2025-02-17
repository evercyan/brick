package xhttp

import (
	"context"
	"fmt"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xurl"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// GET ...
func GET(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// CrawlPage ...
func CrawlPage(url string) (string, error) {
	client := New(WithRequestTimeout(time.Second * 30))
	resp, err := client.Get(context.Background(), url, http.Header{
		HeaderKeyUserAgent: []string{GetUserAgent()},
	})
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// Crawl ...
func Crawl(url string, caches ...bool) (string, error) {
	cache := xlodash.First(caches, true)
	upath := xurl.FullName(url, false)
	fpath := path.Join(os.TempDir(), upath, "index.html")
	if cache {
		if upath == "" {
			return "", fmt.Errorf("invalid url path")
		}
		if xfile.IsExist(fpath) {
			return xfile.Read(fpath), nil
		}
	}
	resp, err := CrawlPage(url)
	if err != nil {
		return "", err
	}
	if cache {
		if err := xfile.Write(fpath, resp); err != nil {
			return "", err
		}
	}
	return resp, nil
}

// Crawl ...
func Crawl2(url string, caches ...bool) (string, error) {
	if !xlodash.First(caches) {
		return CrawlPage(url)
	}
	upath := xurl.Path(url)
	if upath == "" {
		return "", fmt.Errorf("invalid url path")
	}
	fpath := path.Join(os.TempDir(), upath)
	if xfile.IsExist(fpath) {
		return xfile.Read(fpath), nil
	}
	resp, err := GET(url)
	if err != nil {
		return "", err
	}
	if err := xfile.Write(fpath, resp); err != nil {
		return "", err
	}
	return resp, nil
}
