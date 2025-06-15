package xgithub

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/evercyan/brick/xencoding"
	"github.com/evercyan/brick/xjson"
	"github.com/evercyan/brick/xlodash"
)

// Get 获取文件
func (t *Client) Get(ctx context.Context, path string) (string, error) {
	resp, err := t.Request(ctx, "GET", t.Api(path), "")
	if err != nil {
		return "", err
	}
	return t.Response(resp)
}

// Update 新增或更新文件
func (t *Client) Update(
	ctx context.Context,
	path string,
	content string,
	messages ...string,
) error {
	message := xlodash.First(messages, DefaultMessage)
	param := map[string]interface{}{
		"message": message,
		"committer": map[string]string{
			"name":  t.Owner,
			"email": t.Email,
		},
		"content": xencoding.Base64Encode(content),
	}
	// 如果存在 sha, 更新文件
	sha := t.Sha(ctx, path)
	if sha != "" {
		param["sha"] = sha
	}
	resp, err := t.Request(ctx, "PUT", t.Api(path), xencoding.JSONEncode(param))
	if err != nil {
		return err
	}
	_, err = t.Response(resp)
	return err
}

// Delete 删除文件
func (t *Client) Delete(ctx context.Context, path string) error {
	sha := t.Sha(ctx, path)
	if sha == "" {
		return errors.New("获取文件 sha 失败")
	}
	param := map[string]interface{}{
		"message": DefaultMessage,
		"sha":     sha,
	}
	resp, err := t.Request(ctx, "DELETE", t.Api(path), xencoding.JSONEncode(param))
	if err != nil {
		return err
	}
	_, err = t.Response(resp)
	return err
}

// GetCdnUrl 获取文件链接
func (t *Client) GetCdnUrl(path string) string {
	return fmt.Sprintf(CdnURL, t.Owner, t.Repo, path)
}

// GetLastVersion 获取应用最后版本号
func (t *Client) GetLastVersion(ctx context.Context) string {
	tagURL := fmt.Sprintf(TagURL, t.Owner, t.Repo)
	resp, err := t.Request(ctx, "GET", tagURL, "")
	if err != nil {
		return ""
	}
	return xjson.New(resp).Index(0).Key("name").ToString()
}

// GetContent ...
func (t *Client) GetContent(ctx context.Context, path string) string {
	resp, err := t.Get(ctx, path)
	if err != nil {
		return ""
	}
	return xencoding.Base64Decode(xjson.New(resp).Key("content").ToString())
}

// Sha ...
func (t *Client) Sha(ctx context.Context, path string) string {
	resp, _ := t.Get(ctx, path)
	return xjson.New(resp).Key("sha").ToString()
}

// Api ...
func (t *Client) Api(path string) string {
	return fmt.Sprintf(ApiURL, t.Owner, t.Repo, path)
}

// Response ...
func (t *Client) Response(resp string) (string, error) {
	message := xjson.New(resp).Key("message").ToString()
	if message != "" {
		return "", errors.New(message)
	}
	return resp, nil
}

// Request ...
func (t *Client) Request(
	ctx context.Context,
	method string,
	url string,
	data string,
) (string, error) {
	body := bytes.NewReader([]byte(data))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", t.AccessToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
