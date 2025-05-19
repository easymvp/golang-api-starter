package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func GenerateDESKey() (string, error) {
	key := make([]byte, 8)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}

func DecodeDESKey(key string) ([]byte, error) {
	return hex.DecodeString(key)
}

func Encrypt(key, text string) (string, error) {
	keyBytes, err := DecodeDESKey(key)
	if err != nil {
		return "", err
	}
	textBytes := []byte(text)

	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	padding := des.BlockSize - len(textBytes)%des.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	textBytes = append(textBytes, padText...)

	ciphertext := make([]byte, len(textBytes))
	encrypter := cipher.NewCBCEncrypter(block, keyBytes[:des.BlockSize])
	encrypter.CryptBlocks(ciphertext, textBytes)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key, encryptedText string) (string, error) {
	keyBytes, err := DecodeDESKey(key)
	if err != nil {
		return "", err
	}
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := des.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, keyBytes[:des.BlockSize])
	decrypted := make([]byte, len(encryptedBytes))
	decrypter.CryptBlocks(decrypted, encryptedBytes)

	// Remove padding
	padding := decrypted[len(decrypted)-1]
	decrypted = decrypted[:len(decrypted)-int(padding)]

	return string(decrypted), nil
}
