package xcrypto

import (
	"strings"
	"testing"

	"github.com/evercyan/brick/xencoding"
	"github.com/stretchr/testify/assert"
)

var (
	aesKey            = "0000111122223333"
	aesText           = "hello world"
	aesEcbEncrypted   = "GKwjbsCos8ozddpA37gbmg=="
	aesMysqlEncrypted = "18ac236ec0a8b3ca3375da40dfb81b9a"
)

func TestAesEcb(t *testing.T) {
	encrypted := AesECBEncrypt(aesText, aesKey)
	assert.NotEmpty(t, encrypted)
	assert.Equal(t, aesEcbEncrypted, xencoding.Base64Encode(encrypted))

	decrypted := AesECBDecrypt(encrypted, aesKey)
	assert.NotEmpty(t, decrypted)
	assert.Equal(t, aesText, decrypted)

	{
		// 30 > 密钥长度 16
		assert.Equal(t, "aZaMTIJWTyJZjB/xn7Ki2Q==", xencoding.Base64Encode(AesECBEncrypt("hello", strings.Repeat("k", 30))))
	}
}

func TestAesCBC(t *testing.T) {
	key := []byte("abcdefghijklmnop")
	iv := []byte("iviviviviviviviv")
	text := []byte("hello world")

	encrypted, err := AesCBCEncrypt(text, key, iv)
	assert.Nil(t, err)
	assert.NotEmpty(t, encrypted)

	decrypted, err := AesCBCDecrypt(encrypted, key, iv)
	assert.Nil(t, err)
	assert.Equal(t, text, decrypted)
}

func TestAesMysql(t *testing.T) {
	encrypted := AesMysqlEncrypt(aesText, aesKey)
	assert.NotEmpty(t, encrypted)
	assert.Equal(t, aesMysqlEncrypted, encrypted)

	decrypted := AesMysqlDecrypt(encrypted, aesKey)
	assert.NotEmpty(t, decrypted)
	assert.Equal(t, aesText, decrypted)
}

func TestCover(t *testing.T) {
	assert.Empty(t, AesECBDecrypt("a", "a"))
	assert.Equal(t, "a", AesMysqlDecrypt("a", "a"))
}
