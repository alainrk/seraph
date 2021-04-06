// Inspired by and thanks to https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func encrypt(keyString string, plaintext string) (encryptedString string) {
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

func decrypt(keyString string, ciphertext string) (decryptedString string) {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(string(ciphertext))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := aesGCM.NonceSize()
	nonce, encrypted := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, encrypted, nil)

	return string(plaintext)
}

// Prompt the output on stdin and returns the clean input string given from the user
func prompt(output string) (inputString string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(output)
	text, err := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n\r")
	if err != nil {
		panic(err)
	}
	return text
}

func main() {
	// passphrase := "helloworld"
	passphrase := prompt("Insert key: ")
	fmt.Printf("Your key: ---%s---\n", passphrase)
	plaintext := "This is a great secret to keep!"

	keyBytes := sha256.Sum256([]byte(passphrase))
	key := hex.EncodeToString(keyBytes[:])

	fmt.Printf("KeyBytes: %x\nKey: %s\n", keyBytes, key)

	ciphertext := encrypt(key, plaintext)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	deciphertext := decrypt(key, ciphertext)
	fmt.Printf("Decrypted: %s\n", deciphertext)
}
