package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// AesEcbEncrypt ...
func AesEcbEncrypt(text, key string) string {
	bt, bk := []byte(text), []byte(key)
	cip, err := aes.NewCipher(aesEcbKey(bk))
	if err != nil {
		return ""
	}
	blockSize := cip.BlockSize()
	length := (len(bt) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, bt)
	pad := byte(len(plain) - len(bt))
	for i := len(bt); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, blockSize; bs <= len(bt); bs, be = bs+blockSize, be+blockSize {
		cip.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return string(encrypted)
}

// AesEcbDecrypt ...
func AesEcbDecrypt(text, key string) string {
	bt, bk := []byte(text), []byte(key)
	cip, err := aes.NewCipher(aesEcbKey(bk))
	if err != nil {
		return ""
	}
	blockSize := cip.BlockSize()
	decrypted := make([]byte, len(bt))
	for bs, be := 0, blockSize; bs < len(bt); bs, be = bs+blockSize, be+blockSize {
		if be > len(bt) {
			return ""
		}
		cip.Decrypt(decrypted[bs:be], bt[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return string(decrypted[:trim])
}

// aesEcbKey ...
func aesEcbKey(key []byte) []byte {
	res := make([]byte, 16)
	copy(res, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			res[j] ^= key[i]
		}
	}
	return res
}

// ----------------------------------------------------------------

// AesCbcEncrypt ...
func AesCbcEncrypt(text, key string) string {
	bt, bk := []byte(text), []byte(key)
	block, err := aes.NewCipher(bk)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	data := aesCbcPKCS7Padding(bt, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, bk[:blockSize])
	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)
	return base64.StdEncoding.EncodeToString(encrypted)
}

// AesCbcDecrypt ...
func AesCbcDecrypt(text string, key string) string {
	bt, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	bk := []byte(key)
	block, err := aes.NewCipher(bk)
	if err != nil {
		return ""
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, bk[:blockSize])
	decrypted := make([]byte, len(bt))
	blockMode.CryptBlocks(decrypted, bt)
	return string(aecCbcPKCS7UnPadding(decrypted))
}

// aesCbcPKCS7Padding ...
func aesCbcPKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

// aecCbcPKCS7UnPadding ...
func aecCbcPKCS7UnPadding(text []byte) []byte {
	length := len(text)
	unpadding := int(text[length-1])
	return text[:(length - unpadding)]
}

// ----------------------------------------------------------------

// AesMysqlEncrypt 同 mysql hex(aes_encrypt('text', 'key'))
func AesMysqlEncrypt(text, key string) string {
	return hex.EncodeToString([]byte(AesEcbEncrypt(text, key)))
}

// AesMysqlDecrypt 同 mysql aes_decrypt(unhex('text'), 'key')
func AesMysqlDecrypt(text, key string) string {
	b, err := hex.DecodeString(text)
	if err != nil {
		return text
	}
	decrypted := AesEcbDecrypt(string(b), key)
	if decrypted == "" {
		return text
	}
	return decrypted
}
