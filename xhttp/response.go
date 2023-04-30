package xhttp

import (
	"net/http"
)

// Response ...
type Response struct {
	*http.Response

	trace *Trace
	body  []byte
}

// Bytes ...
func (t *Response) Bytes() []byte {
	return t.body
}

// String ...
func (t *Response) String() string {
	return string(t.body)
}

// Trace ...
func (t *Response) Trace() *Trace {
	if t.trace == nil {
		return nil
	}
	t.trace.DNSLookup = t.trace.dnsDone.Sub(t.trace.dnsStart)
	t.trace.TLSHandshake = t.trace.tlsHandshakeDone.Sub(t.trace.tlsHandshakeStart)
	t.trace.ServerTime = t.trace.gotFirstResponseByte.Sub(t.trace.gotConn)
	t.trace.IsConnReused = t.trace.gotConnInfo.Reused
	t.trace.IsConnWasIdle = t.trace.gotConnInfo.WasIdle
	t.trace.ConnIdleTime = t.trace.gotConnInfo.IdleTime

	if t.trace.IsConnReused {
		t.trace.TotalTime = t.trace.endTime.Sub(t.trace.getConn)
	} else {
		t.trace.TotalTime = t.trace.endTime.Sub(t.trace.dnsStart)
	}
	if !t.trace.connectDone.IsZero() {
		t.trace.TCPConnTime = t.trace.connectDone.Sub(t.trace.dnsDone)
	}
	if !t.trace.gotConn.IsZero() {
		t.trace.ConnTime = t.trace.gotConn.Sub(t.trace.getConn)
	}
	if !t.trace.gotFirstResponseByte.IsZero() {
		t.trace.ResponseTime = t.trace.endTime.Sub(t.trace.gotFirstResponseByte)
	}
	if t.trace.gotConnInfo.Conn != nil {
		t.trace.RemoteAddr = t.trace.gotConnInfo.Conn.RemoteAddr()
	}
	return t.trace
}
