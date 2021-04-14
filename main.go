package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

const vaultDirectory = "./vaults/"
const SecretFile = "./vaults/secret.nrk"

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

func promptForStart() (string, error) {
	prompt := promptui.Select{
		Label: "Choose",
		Items: []string{"Open Vault", "New Vault"},
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

func promptForVaults(vaults []string) string {
	prompt := promptui.Select{
		Label: "Choose",
		Items: vaults,
	}

	_, result, err := prompt.Run()
	check(err)

	return result
}

func getVaults() []string {
	var files []string

	err := filepath.Walk(vaultDirectory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.HasPrefix(info.Name(), ".") {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func main() {
	var passphrase string
	var ciphertext string
	var deciphertext string

	// TODO: Non-interactive handling
	// flags := getFlags()

	mode, _ := promptForStart()
	vaults := getVaults()

	chosenVault := promptForVaults(vaults)
	fmt.Println(chosenVault)

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
