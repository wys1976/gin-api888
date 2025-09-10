package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// 密钥长度必须是16(AES-128)、24(AES-192)或32(AES-256)字节
var SecretKey = []byte("wys888wys888wys8") // 16 bytes for AES-128

// PKCS7Padding 进行PKCS7填充[2](@ref)
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding 去除PKCS7填充[2](@ref)
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("empty data")
	}
	unPadding := int(data[length-1])
	if unPadding > length || unPadding > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unPadding)], nil
}

// Encrypt 加密函数 (修复版)[2,5](@ref)
func Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	plaintextBytes := []byte(plaintext)
	// 对明文进行PKCS7填充
	plaintextBytesPadded := PKCS7Padding(plaintextBytes, block.BlockSize())

	// 生成随机IV（16字节）[2](@ref)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %v", err)
	}

	// 创建CBC加密模式
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密填充后的明文
	ciphertext := make([]byte, len(plaintextBytesPadded))
	mode.CryptBlocks(ciphertext, plaintextBytesPadded)

	// 组合IV和密文[2](@ref)
	ivAndCiphertext := make([]byte, len(iv)+len(ciphertext))
	copy(ivAndCiphertext[:aes.BlockSize], iv)
	copy(ivAndCiphertext[aes.BlockSize:], ciphertext)

	return base64.StdEncoding.EncodeToString(ivAndCiphertext), nil
}

// Decrypt 解密函数 (修复版)[2,5](@ref)
func Decrypt(encrypted string) (string, error) {
	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Base64解码
	encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %v", err)
	}

	if len(encryptedBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// 提取IV和实际密文[2](@ref)
	iv := encryptedBytes[:aes.BlockSize]
	ciphertext := encryptedBytes[aes.BlockSize:]

	// 检查密文长度是否为块大小的倍数
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	// 创建CBC解密模式
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密操作
	decryptedBytes := make([]byte, len(ciphertext))
	mode.CryptBlocks(decryptedBytes, ciphertext)

	// 去除PKCS7填充
	unpaddedBytes, err := PKCS7UnPadding(decryptedBytes)
	if err != nil {
		return "", fmt.Errorf("failed to unpad: %v", err)
	}

	return string(unpaddedBytes), nil
}

// EncryptPhone 加密手机号（专为手机号设计）
func EncryptPhone(phone string) (string, error) {
	if phone == "" {
		return "", nil
	}
	return Encrypt(phone)
}

// DecryptPhone 解密手机号
func DecryptPhone(encryptedPhone string) (string, error) {
	if encryptedPhone == "" {
		return "", nil
	}
	return Decrypt(encryptedPhone)
}
