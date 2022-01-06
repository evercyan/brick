package xcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"github.com/spaolacci/murmur3"
)

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
