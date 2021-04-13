package main

import (
	"errors"
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

func promptForPassword(label string, validate func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	return result, err
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

	// TODO: Non-interactive handling
	// flags := getFlags()

	mode, err := promptForMode()
	check(err)

	// Check if secret file exists, create otherwise
	if _, err := os.Stat(SecretFile); err != nil {
		f, err := os.Create(SecretFile)
		check(err)
		defer f.Close()
	}

	if mode == "Encrypt" {
		passphrase, _ = promptForPassword("Password", validatePassword)
		validatorConfirm := func(s string) error {
			err := validatePassword(s)
			if err != nil {
				return err
			}
			if s != passphrase {
				return errors.New(passwordsNotMatchingError)
			}
			return nil
		}
		_, _ = promptForPassword("Confirm", validatorConfirm)

		key := hashPassphrase(passphrase)

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

		passphrase, _ = promptForPassword("Password", validatePassword)
		key := hashPassphrase(passphrase)

		deciphertext, err = decrypt(key, ciphertext)
		if err != nil {
			fmt.Println("Error decoding crypted data. Check your password.", err)
			return
		}
		fmt.Printf("Decrypted: %s\n", deciphertext)
	}
}
