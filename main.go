package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	dateTimeFormat string = "2006-01-02T15:04:05-0700"
)

const (
	vaultDirectory          = "./vaults/"
	SecretFile              = "./vaults/secret.nrk"
	vaultAlreadyExistsError = "vault already exists"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func openVaultHandling() {
	vaults := getVaults()
	_, chosenVault, _ := promptForSelect("Choose", vaults)

	fmt.Println("Opening vault", chosenVault, "...")

	vaultPath := filepath.Join(vaultDirectory, chosenVault)

	dat, _ := ioutil.ReadFile(vaultPath)
	ciphertext := string(dat)

	passphrase, _ := promptForPassword("Password", validatePassword)
	key := hashPassphrase(passphrase)

	deciphertext, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Error decoding crypted data. Check your password.", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", deciphertext)
}

func newVaultHandling() {
	// Validate already exist vault
	vaults := getVaults()
	validatorVaultNotExists := func(s string) error {
		for _, vault := range vaults {
			if s+".vault" == vault {
				return errors.New(vaultAlreadyExistsError)
			}
		}
		return nil
	}

	newVaultName := promptForTextValid("Vault name", validatorVaultNotExists)
	fmt.Println("new vault", newVaultName)

	newVaultPath := filepath.Join(vaultDirectory, newVaultName+".vault")

	// Create the vault
	f, err := os.Create(newVaultPath)
	check(err)
	defer f.Close()

	// Init the vault's password
	passphrase, _ := promptForPassword("Password", validatePassword)
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

	v := newVaultEmpty()

	plaintext := v.marshal()
	ciphertext := encrypt(key, plaintext)
	err = ioutil.WriteFile(newVaultPath, []byte(ciphertext), 0644)
	check(err)
	fmt.Printf("Encrypted file to %s\n", SecretFile)
}

func main() {
	// TODO: Non-interactive handling
	// flags := getFlags()

	const (
		openVault int = iota
		newVault
	)
	index, _, _ := promptForSelect("Choose", []string{"Open Vault", "New Vault"})

	// Opening vault
	if index == openVault {
		openVaultHandling()
	} else if index == newVault {
		newVaultHandling()
	}
}
