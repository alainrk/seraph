package main

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := hashPassphrase("passphrase")
	plaintext := "This is a test"
	ciphertext := encrypt(key, plaintext)
	deciphertext := decrypt(key, ciphertext)

	if deciphertext != plaintext {
		t.Errorf("Error on crypt/decrypt. Expected \"%s\", given \"%s\"", plaintext, deciphertext)
	}
}
