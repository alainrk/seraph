package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

func chooseVault() (*vault, error) {
	vaults := getVaults()
	_, chosenVault, _ := promptForSelect("Choose", vaults)

	fmt.Println("Opening vault", chosenVault, "...")

	vaultPath := filepath.Join(vaultDirectory, chosenVault)

	dat, _ := ioutil.ReadFile(vaultPath)
	ciphertext := string(dat)

	passphrase, _ := promptForPassword("Password", validatePassword)
	key := hashPassphrase(passphrase)

	vaultMarshaled, err := decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Error decoding crypted data. Check your password.", err)
		return nil, err
	}

	vault := newVaultEmpty()
	vault.unmarshal(vaultMarshaled)
	return vault, nil
}

func insertSecretHandling(vault vault) {
	// TODO Test
	s := secret{}
	s.Name = "Lorem"
	s.Username = "ipsum"
	s.Email = "dolor@s.it"
	s.Password = "amet"
	s.ApiKey = "0398509234"
	s.Notes = "Test 1"
	s.CreatedAt = time.Now().Format(dateTimeFormat)
	vault.add(s)

	// TODO --- I need a context, created at main that's been passed around
	// I need to keep the password in some way and other data
	//
	// marshaled := vault.marshal()
	// ciphertext := encrypt(key, marshaled)
	// err = ioutil.WriteFile(newVaultPath, []byte(ciphertext), 0644)
	// check(err)
	// fmt.Printf("Encrypted file to %s\n", SecretFile)
}

func getSecretHandling(vault vault) {
	keys := make([]string, 0)
	for k, _ := range vault.KeysMap {
		keys = append(keys, k)
	}

	_, key, _ := promptForSelect("Choose", keys)
	fmt.Println(vault.KeysMap[key])
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
