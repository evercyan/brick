package xcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// AesECBEncrypt ...
func AesECBEncrypt(text, key string) string {
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

// AesECBDecrypt ...
func AesECBDecrypt(text, key string) string {
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

// PKCS7Padding ...
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding ...
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesCBCEncrypt AES-128-CBC 加密
func AesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data = PKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return crypted, nil
}

// AesCBCDecrypt ase-128-cbc 解密
func AesCBCDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	data := make([]byte, len(crypted))
	blockMode.CryptBlocks(data, crypted)
	data = PKCS7UnPadding(data)
	return data, nil
}

// ----------------------------------------------------------------

// AesMysqlEncrypt 同 mysql hex(aes_encrypt('text', 'key'))
func AesMysqlEncrypt(text, key string) string {
	return hex.EncodeToString([]byte(AesECBEncrypt(text, key)))
}

// AesMysqlDecrypt 同 mysql aes_decrypt(unhex('text'), 'key')
func AesMysqlDecrypt(text, key string) string {
	b, err := hex.DecodeString(text)
	if err != nil {
		return text
	}
	decrypted := AesECBDecrypt(string(b), key)
	if decrypted == "" {
		return text
	}
	return decrypted
}
