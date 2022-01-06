package xhttp

import (
	"crypto/tls"
	"time"
)

// Option ...
type Option struct {
	SSLEnabled            bool
	TLSConfig             *tls.Config
	Compressed            bool
	HandshakeTimeout      time.Duration
	ResponseHeaderTimeout time.Duration
	RequestTimeout        time.Duration
	ConnsPerHost          int
	RetryTimes            int
}

// ----------------------------------------------------------------

// defaultOption ...
var defaultOption = &Option{
	Compressed:            true,
	HandshakeTimeout:      30 * time.Second,
	ResponseHeaderTimeout: 30 * time.Second,
	RequestTimeout:        30 * time.Second,
	ConnsPerHost:          5,
	RetryTimes:            3,
}

// setOptionDefaultValue ...
func setOptionDefaultValue(option *Option) *Option {
	if option == nil {
		return defaultOption
	}
	if option.RequestTimeout <= 0 {
		option.RequestTimeout = defaultOption.RequestTimeout
	}
	if option.HandshakeTimeout <= 0 {
		option.HandshakeTimeout = defaultOption.HandshakeTimeout
	}
	if option.ResponseHeaderTimeout <= 0 {
		option.ResponseHeaderTimeout = defaultOption.ResponseHeaderTimeout
	}
	if option.ConnsPerHost <= 0 {
		option.ConnsPerHost = defaultOption.ConnsPerHost
	}
	if option.RetryTimes <= 0 {
		option.RetryTimes = defaultOption.RetryTimes
	}
	return option
}
