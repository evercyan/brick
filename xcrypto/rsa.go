package xcrypto

import (
	"github.com/wenzhenxi/gorsa"
)

// RsaPubEncrypt 公钥加密
func RsaPubEncrypt(text, key string) (string, error) {
	return gorsa.PublicEncrypt(text, key)
}

// RsaPubDecrypt 公钥解密
func RsaPubDecrypt(text, key string) (string, error) {
	return gorsa.PublicDecrypt(text, key)
}

// RsaPriEncrypt 私钥加密
func RsaPriEncrypt(text, key string) (string, error) {
	return gorsa.PriKeyEncrypt(text, key)
}

// RsaPriDecrypt 私钥解密
func RsaPriDecrypt(text, key string) (string, error) {
	return gorsa.PriKeyDecrypt(text, key)
}
