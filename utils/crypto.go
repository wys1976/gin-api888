package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var (
	secretKey = []byte("wys888")
	iv        = []byte("1234567812345678")
)

func Encrypt(plaintext string) string {
	block, _ := aes.NewCipher(secretKey)
	plaintextBytes := []byte(plaintext)

	ciphertext := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintextBytes)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(encrypted string) string {
	block, _ := aes.NewCipher(secretKey)
	ciphertext, _ := base64.StdEncoding.DecodeString(encrypted)

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	return string(plaintext)
}
