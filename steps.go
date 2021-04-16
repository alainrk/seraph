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

// Step 1
const (
	back int = iota
	getSecret
	insertSecret
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

func chooseVault(ctx *Context) error {
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

func openedVaultHandling(ctx *Context) {
	for {
		index, _, _ := promptForSelect("Choose", []string{"Back", "Get secret", "Insert secret"})

		if index == back {
			return
		} else if index == insertSecret {
			insertSecretHandling(ctx)
		} else if index == getSecret {
			getSecretHandling(ctx)
		}
	}
}

func insertSecretHandling(ctx *Context) {
	var choice string
	var value string
	changed := false

	fields := []string{"Exit", "Name", "Username", "Email", "Password", "ApiKey", "Notes"}
	s := secret{}

	for {
		_, choice, _ = promptForSelect("Choose a field to edit or exit", fields)
		if choice == "Exit" {
			break
		}
		changed = true
		value, _ = promptForText(choice)
		switch choice {
		case "Name":
			s.Name = value
		case "Username":
			s.Username = value
		case "Email":
			s.Email = value
		case "Password":
			s.Password = value
		case "ApiKey":
			s.ApiKey = value
		case "Notes":
			s.Notes = value
		}
	}

	if !changed {
		return
	}

	s.CreatedAt = time.Now().Format(dateTimeFormat)
	s.UpdatedAt = time.Now().Format(dateTimeFormat)

	ctx.vault.add(s)
	saveVault(ctx, true)
}

func getSecretHandling(ctx *Context) {
	keys := make([]string, 0)
	for k, _ := range ctx.vault.KeysMap {
		keys = append(keys, k)
	}

	_, key, _ := promptForSelect("Choose", keys)

	fmt.Println("\nName:", ctx.vault.KeysMap[key].Name)
	fmt.Println("Username:", ctx.vault.KeysMap[key].Username)
	fmt.Println("Email:", ctx.vault.KeysMap[key].Email)
	fmt.Println("Password:", ctx.vault.KeysMap[key].Password)
	fmt.Println("ApiKey:", ctx.vault.KeysMap[key].ApiKey)
	fmt.Println("Notes:", ctx.vault.KeysMap[key].Notes+"\n")

	promptToJustWait()
}

func newVaultHandling(ctx *Context) {
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

	saveVault(ctx, false)
}

func saveVault(ctx *Context, askConfirm bool) {
	if askConfirm {
		confirm, _ := promptForConfirm("Save")
		if confirm == false {
			fmt.Println("Changes not saved")
			return
		}
	}
	marshaledPlainText := ctx.vault.marshal()
	marshaledCipherText := encrypt(ctx.hashedPassword, marshaledPlainText)
	err := ioutil.WriteFile(ctx.vault.path, []byte(marshaledCipherText), 0644)
	check(err)
	fmt.Printf("Encrypted file to %s\n", ctx.vault.path)
}
