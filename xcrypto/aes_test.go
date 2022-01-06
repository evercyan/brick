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
	aesCbcEncrypted   = "UnRNRit0Nm80eUgxd0szdVZIMzRPUT09"
	aesMysqlEncrypted = "18ac236ec0a8b3ca3375da40dfb81b9a"
)

func TestAesEcb(t *testing.T) {
	encrypted := AesEcbEncrypt(aesText, aesKey)
	assert.NotEmpty(t, encrypted)
	assert.Equal(t, aesEcbEncrypted, xencoding.Base64Encode(encrypted))

	decrypted := AesEcbDecrypt(encrypted, aesKey)
	assert.NotEmpty(t, decrypted)
	assert.Equal(t, aesText, decrypted)

	{
		// 30 > 密钥长度 16
		assert.Equal(t, "aZaMTIJWTyJZjB/xn7Ki2Q==", xencoding.Base64Encode(AesEcbEncrypt("hello", strings.Repeat("k", 30))))
	}
}

func TestAesCbc(t *testing.T) {
	encrypted := AesCbcEncrypt(aesText, aesKey)
	assert.NotEmpty(t, encrypted)
	assert.Equal(t, aesCbcEncrypted, xencoding.Base64Encode(encrypted))

	decrypted := AesCbcDecrypt(encrypted, aesKey)
	assert.NotEmpty(t, decrypted)
	assert.Equal(t, aesText, decrypted)
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
	assert.Empty(t, AesEcbDecrypt("a", "a"))
	assert.Empty(t, AesCbcDecrypt("a", "a"))
	assert.Equal(t, "a", AesMysqlDecrypt("a", "a"))
}
