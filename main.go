// Inspired by and thanks to https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/

package main

import (
	"fmt"
)

func main() {
	getFlags()

	// passphrase := "helloworld"
	passphrase := prompt("Insert key: ")
	fmt.Printf("Your key: ---%s---\n", passphrase)

	plaintext := "This is a great secret to keep!"

	key := hashPassphrase(passphrase)

	ciphertext := encrypt(key, plaintext)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	deciphertext := decrypt(key, ciphertext)
	fmt.Printf("Decrypted: %s\n", deciphertext)
}
