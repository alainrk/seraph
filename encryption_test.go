package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	key := hashPassphrase("passphrase")
	plaintext := "This is a test"
	ciphertext := encrypt(key, plaintext)
	deciphertext, _ := decrypt(key, ciphertext)
	assert.Equal(t, plaintext, deciphertext, "they should be equal")
}
