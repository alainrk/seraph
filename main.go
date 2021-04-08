package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const SecretFile = "assets/secret.nrk"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flags := getFlags()
	var passphrase string
	var ciphertext string
	var deciphertext string

	if flags.passphrase == "" {
		passphrase = prompt("Insert key: ")
	} else {
		passphrase = flags.passphrase
	}

	key := hashPassphrase(passphrase)

	// Check secret file exists, create otherwise
	if _, err := os.Stat(SecretFile); err != nil {
		f, err := os.Create(SecretFile)
		check(err)
		// close immediately here
		f.Close()
	}

	if flags.encryptMode {
		plaintext := prompt("Insert secret: ")
		ciphertext = encrypt(key, plaintext)
		err := ioutil.WriteFile(SecretFile, []byte(ciphertext), 0644)
		check(err)
		fmt.Printf("Encrypted file to %s\n", SecretFile)
	}

	if flags.decryptMode {
		dat, err := ioutil.ReadFile(SecretFile)
		check(err)
		ciphertext = string(dat)
		deciphertext = decrypt(key, ciphertext)
		fmt.Printf("Decrypted: %s\n", deciphertext)
	}

}
