package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

const SecretFile = "assets/secret.nrk"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func promptForPassword(validate func(string) error) string {
	prompt := promptui.Prompt{
		Label:    "Passphrase",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	check(err)

	return result
}

func promptForMode() (string, error) {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"Encrypt", "Decrypt"},
	}

	_, mode, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return mode, nil
}

func promptForText(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	check(err)

	return result
}

func main() {
	var passphrase string
	var ciphertext string
	var deciphertext string

	mode, err := promptForMode()
	check(err)

	var passwdValidator func(string) error
	if mode == "Encrypt" {
		passwdValidator = validatePassword
	} else {
		passwdValidator = validateAlwaysString
	}

	passphrase = promptForPassword(passwdValidator)
	key := hashPassphrase(passphrase)

	// Check if secret file exists, create otherwise
	if _, err := os.Stat(SecretFile); err != nil {
		f, err := os.Create(SecretFile)
		defer f.Close()
		check(err)
	}

	if mode == "Encrypt" {
		plaintext := promptForText("Insert secret: ")
		ciphertext = encrypt(key, plaintext)
		err := ioutil.WriteFile(SecretFile, []byte(ciphertext), 0644)
		check(err)
		fmt.Printf("Encrypted file to %s\n", SecretFile)
	}

	if mode == "Decrypt" {
		dat, err := ioutil.ReadFile(SecretFile)
		check(err)
		ciphertext = string(dat)
		deciphertext = decrypt(key, ciphertext)
		fmt.Printf("Decrypted: %s\n", deciphertext)
	}
}

// OLD modality -> maybe to consider for non interactive mode
func flagsHandling() {
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
