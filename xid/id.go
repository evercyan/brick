package xid

import (
	"crypto/rand"
	"io"

	"github.com/evercyan/brick/xcrypto"
	"github.com/evercyan/brick/xencoding"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
)

// GUID ...
func GUID() string {
	b := make([]byte, 48)
	io.ReadFull(rand.Reader, b)
	return xcrypto.Md5(xencoding.Base64Encode(string(b)))
}

// UUID ...
func UUID() string {
	return uuid.NewV4().String()
}

// XID ...
func XID() string {
	return xid.New().String()
}
