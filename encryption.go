// Inspired by and thanks to https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func encrypt(keyString string, plaintext string) string {
	key, _ := hex.DecodeString(keyString)
	plaintextBytes := []byte(plaintext)

	// key[:] => shortcut to copy and get a slice of []bytes
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// see: https://en.wikipedia.org/wiki/Galois/Counter_Mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	// nonce creation from gcm
	nonce := make([]byte, aesGCM.NonceSize())

	// encryption of data with gcm seal
	ciphertext := aesGCM.Seal(nonce, nonce, plaintextBytes, nil)

	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(keyString string, ciphertext string) (string, error) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(string(ciphertext))

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, encrypted := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func hashPassword(password string) string {
	keyBytes := sha256.Sum256([]byte(password))
	key := hex.EncodeToString(keyBytes[:])
	return key
}
