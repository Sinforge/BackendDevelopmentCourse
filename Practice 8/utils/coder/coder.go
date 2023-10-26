package coder

import (
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"

)

var (
	key = []byte("supersecretkey12")
)



func DecryptCookie(encrypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	iv := encrypted[:gcm.NonceSize()]
	ciphertext := encrypted[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
	
func EncryptCookie(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, iv, plaintext, nil)

	encrypted := append(iv, ciphertext...)

	return encrypted, nil
}
