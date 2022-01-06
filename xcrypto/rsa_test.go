package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 私钥: openssl genrsa -out rsa_private_key.pem 1024
	rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCwVZrnmWdYeK2QPkCw+GbX5oa7PyiHVlvgzwqjyiNDgzF63kWI
5MkqUfW5QEcPFrlRPfvwGlYUdzEDOajbvHmsdWFdnpY4pfkUM5XKVNCIxeZ4Eeo+
hwt9XsVVFt6v9aVP8BivVvpvKSCxUwF8DyhVhenLmuWmhi4BK+cvTyJfDQIDAQAB
AoGAbJksP7qghUI9UnqkiNlTLWCSVbu+ECvFhhy85sbVFP01egBuPrL4mZHEjgTi
Po62LyzHfsYZgZ0umFfogPwAyonSa2EujDrlcWanfJBp9Veovf0QoN6KIQnntxe0
d5YjYgo+0uqnAyMYICNL2t+fTjJPpxMrufB7DaK2PRYaRsECQQDcxyKGT2EwDOVa
3RtCa1fXn4DZZip31gvwhZiMvwQXS5idOfTfpBLWKVTjL7EW9RnXPOjHVXeCYX4C
ptfX5vUZAkEAzHdWw4cB0e4vdHB1lUpgTLsM2aovY3ueyu9YE5IVaxLrO+EEXFs7
q/UlYv6sCJ0XnuiZYvbdAPN7BMt8pNbkFQJBAM7jdUDzhhmnHA7IAFF/kfOnrvEK
wmVGGi4so0XRgp3p43wC4avpbxVt6fRzrrnauXpvw5t4RePSRGlru/zAm5ECQEtc
NOtuKDqS2oTFKmFy/1lom8ziEANPvfA4FTNpZWGIoJD6V5weuDih6zy4dvnZxKn6
OwahzEUceJwE0BUFax0CQQCqDVS96526EFPP3EhIxLGq5iPvJBZRiVI1plzQgsv7
BuWKdSqiQM4rvxcPbkSPZPUj+BqrDxvAq8hYd0edmQiK
-----END RSA PRIVATE KEY-----`
	// 公钥: openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
	rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCwVZrnmWdYeK2QPkCw+GbX5oa7
PyiHVlvgzwqjyiNDgzF63kWI5MkqUfW5QEcPFrlRPfvwGlYUdzEDOajbvHmsdWFd
npY4pfkUM5XKVNCIxeZ4Eeo+hwt9XsVVFt6v9aVP8BivVvpvKSCxUwF8DyhVhenL
muWmhi4BK+cvTyJfDQIDAQAB
-----END PUBLIC KEY-----`
	rsaText = "hello world"
)

func TestRsa(t *testing.T) {
	{
		// 客户端公钥加密, 服务端私钥解密
		encrypted, _ := RsaPubEncrypt(rsaText, rsaPublicKey)
		assert.NotEmpty(t, encrypted)
		decrypted, _ := RsaPriDecrypt(encrypted, rsaPrivateKey)
		assert.NotEmpty(t, decrypted)
		assert.Equal(t, rsaText, decrypted)
	}
	{
		// 服务端私钥加密, 客户端公钥解密
		encrypted, _ := RsaPriEncrypt(rsaText, rsaPrivateKey)
		assert.NotEmpty(t, encrypted)
		decrypted, _ := RsaPubDecrypt(encrypted, rsaPublicKey)
		assert.NotEmpty(t, decrypted)
		assert.Equal(t, rsaText, decrypted)
	}
}
