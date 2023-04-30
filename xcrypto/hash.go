package xcrypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/spaolacci/murmur3"
)

// Md5 ...
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Hash ...
func Hash(str string, seed int) uint64 {
	dataSha := sha256.Sum256([]byte(str))
	data := dataSha[:]
	m := murmur3.New64WithSeed(uint32(seed))
	m.Write(data)
	return m.Sum64()
}

// HmacSha256 ...
func HmacSha256(src string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(src))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Sha1 ...
func Sha1(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256 ...
func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
