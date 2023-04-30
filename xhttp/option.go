package xhttp

import (
	"net"
	"net/http"
	"time"
)

// Option ...
type Option struct {
	RetryTimes     int             // 请求重试次数
	TraceEnable    bool            // Trace 开关
	RequestTimeout time.Duration   // 请求超时
	Dialer         *net.Dialer     // dialer 配置
	Transport      *http.Transport // transport 配置
}

// ----------------------------------------------------------------

// OptionFn ...
type OptionFn func(*Option)

// WithTrace 可以通过返回 Response.Trace 查询
func WithTraceEnable() OptionFn {
	return func(o *Option) {
		o.TraceEnable = true
	}
}

// WithRetryTimes 重试次数
func WithRetryTimes(v int) OptionFn {
	return func(o *Option) {
		if v > 0 {
			o.RetryTimes = v
		}
	}
}

// WithRequestTimeout 请求超时
func WithRequestTimeout(v time.Duration) OptionFn {
	return func(o *Option) {
		o.RequestTimeout = v
	}
}

// WithDialer ...
func WithDialer(v *net.Dialer) OptionFn {
	return func(o *Option) {
		o.Dialer = v
	}
}

// WithTransport ...
func WithTransport(v *http.Transport) OptionFn {
	return func(o *Option) {
		o.Transport = v
	}
}

// ----------------------------------------------------------------

// defaultOption ...
var defaultOption *Option

func init() {
	defaultOption = &Option{
		RetryTimes:     1,
		RequestTimeout: 60 * time.Second,
		Dialer: &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}
