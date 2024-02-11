package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(unEncryptedData string, key string) string {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	// convert string to bytes
	plaintext := []byte(unEncryptedData)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)

	return Encode(cipherText)
}

func Decrypt(encryptedData string, key string) string {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	// convert string to bytes
	cipherText := Decode(encryptedData)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText)
}

func Tokenize(data string) string {
	plainText := []byte(data)
	hash := sha1.Sum(plainText)
	return hex.EncodeToString(hash[:])
}
