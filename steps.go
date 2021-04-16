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

func chooseVault(ctx Context) error {
	vaults := getVaults()
	_, vaultName, _ := promptForSelect("Choose", vaults)

	fmt.Println("Opening vault", vaultName, "...")

	vaultPath := filepath.Join(vaultDirectory, vaultName)

	dat, _ := ioutil.ReadFile(vaultPath)
	ciphertext := string(dat)

	password, _ := promptForPassword("Password", validatePassword)
	ctx.hashedPassword = hashPassword(password)

	vaultMarshaled, err := decrypt(ctx.hashedPassword, ciphertext)
	if err != nil {
		fmt.Println("Error decoding crypted data. Check your password.", err)
		return err
	}

	vault := newVaultEmpty()
	vault.name = vaultName
	vault.path = vaultPath
	vault.unmarshal(vaultMarshaled)

	ctx.vault = vault
	return nil
}

func insertSecretHandling(ctx Context) {
	// TODO Test
	s := secret{}
	s.Name = "Lorem"
	s.Username = "ipsum"
	s.Email = "dolor@s.it"
	s.Password = "amet"
	s.ApiKey = "0398509234"
	s.Notes = "Test 1"
	s.CreatedAt = time.Now().Format(dateTimeFormat)
	ctx.vault.add(s)

	// Wanna marshal here already? Or maybe wait for save? Or ask the user?
	saveVault(ctx)
}

func getSecretHandling(ctx Context) {
	keys := make([]string, 0)
	for k, _ := range ctx.vault.KeysMap {
		keys = append(keys, k)
	}

	_, key, _ := promptForSelect("Choose", keys)
	fmt.Println(ctx.vault.KeysMap[key])
}

func newVaultHandling(ctx Context) {
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
	newVaultName = newVaultName + ".vault"
	fmt.Println("new vault", newVaultName)

	newVaultPath := filepath.Join(vaultDirectory, newVaultName)

	// Create the vault
	f, err := os.Create(newVaultPath)
	check(err)
	defer f.Close()

	// Init the vault's password
	password, _ := promptForPassword("Password", validatePassword)
	validatorConfirm := func(s string) error {
		err := validatePassword(s)
		if err != nil {
			return err
		}
		if s != password {
			return errors.New(passwordsNotMatchingError)
		}
		return nil
	}
	_, _ = promptForPassword("Confirm", validatorConfirm)

	ctx.hashedPassword = hashPassword(password)

	v := newVaultEmpty()
	v.name = newVaultName
	v.path = newVaultPath

	ctx.vault = v
	saveVault(ctx)
}

func saveVault(ctx Context) {
	marshaledPlainText := ctx.vault.marshal()
	marshaledCipherText := encrypt(ctx.hashedPassword, marshaledPlainText)
	err := ioutil.WriteFile(ctx.vault.path, []byte(marshaledCipherText), 0644)
	check(err)
	fmt.Printf("Encrypted file to %s\n", ctx.vault.path)
}
