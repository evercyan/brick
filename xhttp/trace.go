package xhttp

import (
	"context"
	"crypto/tls"
	"net"
	"net/http/httptrace"
	"time"
)

// Trace ...
type Trace struct {
	getConn              time.Time
	dnsStart             time.Time
	dnsDone              time.Time
	connectDone          time.Time
	tlsHandshakeStart    time.Time
	tlsHandshakeDone     time.Time
	gotConn              time.Time
	gotFirstResponseByte time.Time
	endTime              time.Time
	gotConnInfo          httptrace.GotConnInfo

	DNSLookup      time.Duration `json:"dns_lookup"`
	ConnTime       time.Duration `json:"conn_time"`
	TCPConnTime    time.Duration `json:"tcp_conn_time"`
	TLSHandshake   time.Duration `json:"tls_handshake"`
	ServerTime     time.Duration `json:"server_time"`
	ResponseTime   time.Duration `json:"response_time"`
	TotalTime      time.Duration `json:"total_time"`
	IsConnReused   bool          `json:"is_conn_reused"`
	IsConnWasIdle  bool          `json:"is_conn_was_idle"`
	ConnIdleTime   time.Duration `json:"conn_idle_time"`
	RequestAttempt int           `json:"request_attempt"`
	RemoteAddr     net.Addr      `json:"remote_addr"`
}

// WithClientTrace ...
func (t *Trace) WithClientTrace(ctx context.Context) context.Context {
	return httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		DNSStart: func(_ httptrace.DNSStartInfo) {
			t.dnsStart = time.Now()
		},
		DNSDone: func(_ httptrace.DNSDoneInfo) {
			t.dnsDone = time.Now()
		},
		ConnectStart: func(_, _ string) {
			if t.dnsDone.IsZero() {
				t.dnsDone = time.Now()
			}
			if t.dnsStart.IsZero() {
				t.dnsStart = t.dnsDone
			}
		},
		ConnectDone: func(net, addr string, err error) {
			t.connectDone = time.Now()
		},
		GetConn: func(_ string) {
			t.getConn = time.Now()
		},
		GotConn: func(info httptrace.GotConnInfo) {
			t.gotConn = time.Now()
			t.gotConnInfo = info
		},
		GotFirstResponseByte: func() {
			t.gotFirstResponseByte = time.Now()
		},
		TLSHandshakeStart: func() {
			t.tlsHandshakeStart = time.Now()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			t.tlsHandshakeDone = time.Now()
		},
	})
}

// Finish ...
func (t *Trace) Finish() {
	t.endTime = time.Now()
}
