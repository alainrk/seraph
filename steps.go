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

func chooseVault(app *Context) error {
	var vaultMarshaled string
	var hashedPassword string
	var err error

	vaults := getVaults()

	if len(vaults) == 0 {
		fmt.Println("You have no vaults")
		promptToJustWait()
		return errors.New("No vaults")
	}

	_, vaultName, _ := promptForSelect("Choose", vaults)

	fmt.Println("Opening vault", vaultName, "...")

	vaultPath := filepath.Join(vaultDirectory, vaultName)

	dat, _ := ioutil.ReadFile(vaultPath)
	ciphertext := string(dat)

	for {
		password, _ := promptForPassword("Password", validatePassword)
		hashedPassword = hashPassword(password)
		vaultMarshaled, err = decrypt(hashedPassword, ciphertext)
		if err != nil {
			fmt.Println("Error decoding crypted data. Check your password.", err)

			retry, _ := promptForConfirm("Retry?")
			if !retry {
				return err
			}
		}
		app.hashedPassword = hashedPassword
		break
	}

	vault := newVaultEmpty()
	vault.name = vaultName
	vault.path = vaultPath
	vault.unmarshal(vaultMarshaled)

	app.vault = vault
	return nil
}

func openedVaultHandling(app *Context) {
	for {
		index, _, _ := promptForSelect("Choose", []string{"Back", "Get secret", "Insert secret"})

		if index == back {
			return
		} else if index == insertSecret {
			insertSecretHandling(app)
		} else if index == getSecret {
			getSecretHandling(app)
		}
	}
}

func insertSecretHandling(app *Context) {
	var choice string
	var value string
	changed := false

	fields := []string{"Exit/Save", "Username", "Email", "Password", "ApiKey", "Notes"}
	s := secret{}

	nameValidator := func(name string) error {
		value = strings.TrimSpace(name)
		if _, ok := app.vault.KeysMap[value]; ok {
			return errors.New("This item already exists, choose another name")
		}
		if len(value) == 0 {
			return errors.New("Enter a non-empty name")
		}
		return nil
	}

	value = promptForTextValid("Choose a name", nameValidator)
	s.Name = value

	for {
		_, choice, _ = promptForSelect("Choose a field to edit or exit", fields)
		if choice == "Exit/Save" {
			break
		}
		value, _ = promptForText(choice)
		switch choice {
		case "Username":
			s.Username = value
			changed = true
		case "Email":
			s.Email = value
			changed = true
		case "Password":
			s.Password = value
			changed = true
		case "ApiKey":
			s.ApiKey = value
			changed = true
		case "Notes":
			s.Notes = value
			changed = true
		}
	}

	if !changed {
		return
	}

	s.CreatedAt = time.Now().Format(dateTimeFormat)
	s.UpdatedAt = time.Now().Format(dateTimeFormat)

	app.vault.add(s)
	saveVault(app, true)
	clearScreen()
}

func getSecretHandling(app *Context) {
	keys := make([]string, 0)
	for k, _ := range app.vault.KeysMap {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		fmt.Println("No items in this vault")
		promptToJustWait()
		clearScreen()
		return
	}

	_, key, _ := promptForSelect("Choose", keys)

	fmt.Printf("\nName: %s", app.vault.KeysMap[key].Name)
	fmt.Printf("\nUsername: %s", app.vault.KeysMap[key].Username)
	fmt.Printf("\nEmail: %s", app.vault.KeysMap[key].Email)
	fmt.Printf("\nPassword: %s", app.vault.KeysMap[key].Password)
	fmt.Printf("\nApiKey: %s", app.vault.KeysMap[key].ApiKey)
	fmt.Printf("\nNotes: %s\n\n", app.vault.KeysMap[key].Notes)

	promptToJustWait()
	clearScreen()
}

func newVaultHandling(app *Context) error {
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
	newVaultPath := filepath.Join(vaultDirectory, newVaultName)

	// Create the vault
	f, err := os.Create(newVaultPath)
	defer f.Close()
	if err != nil {
		return err
	}

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
	clearScreen()

	app.hashedPassword = hashPassword(password)

	v := newVaultEmpty()
	v.name = newVaultName
	v.path = newVaultPath
	app.vault = v

	saveVault(app, false)
	return nil
}

func saveVault(app *Context, askConfirm bool) {
	if askConfirm {
		confirm, _ := promptForConfirm("Save")
		if confirm == false {
			fmt.Println("Changes not saved")
			return
		}
	}
	marshaledPlainText := app.vault.marshal()
	marshaledCipherText := encrypt(app.hashedPassword, marshaledPlainText)
	err := ioutil.WriteFile(app.vault.path, []byte(marshaledCipherText), 0644)
	check(err)
	fmt.Printf("Encrypted file to %s\n", app.vault.path)
}
