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
	editSecret
	deleteSecret
)

func initVaultDirectory() error {
	if _, err := os.Stat(vaultDirectory); os.IsNotExist(err) {
		err = os.Mkdir(vaultDirectory, 0755)
		if err != nil {
			check(err)
		}
	}
	return nil
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

func chooseVault(app *Context) error {
	var vaultMarshaled string
	var hashedPassword string
	var err error

	vaults := getVaults()

	if len(vaults) == 0 {
		fmt.Println("You have no vaults")
		promptToJustWait()
		return errors.New("no vaults")
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

			retry, _ := promptForConfirm("Retry")
			if !retry {
				return err
			}
			continue
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
		index, _, _ := promptForSelect("Choose", []string{"Back", "Get secret", "Insert secret", "Edit secret", "Delete secret"})

		if index == back {
			return
		} else if index == insertSecret {
			insertSecretHandling(app)
		} else if index == getSecret {
			getSecretHandling(app)
		} else if index == editSecret {
			editSecretHandling(app)
		} else if index == deleteSecret {
			deleteSecretHandling(app)
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
			return errors.New("this item already exists, choose another name")
		}
		if len(value) == 0 {
			return errors.New("enter a non-empty name")
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
		field, _ := s.assignValueToSecretStringField(choice, value)
		if field == choice {
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

func editSecretHandling(app *Context) {
	var choice string
	var value string
	changed := false

	keys := make([]string, 0)
	for k := range app.vault.KeysMap {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		fmt.Println("No items in this vault")
		promptToJustWait()
		clearScreen()
		return
	}

	_, key, _ := promptForSelect("Choose", keys)

	currentSecret := app.vault.KeysMap[key]

	tmpSecret := secret{}
	tmpSecret.deepCopy(currentSecret)

	fields := []string{"Exit/Save", "Username", "Email", "Password", "ApiKey", "Notes"}

	for {
		_, choice, _ = promptForSelect("Choose a field to edit or exit", fields)
		if choice == "Exit/Save" {
			break
		}

		promptText := choice
		currentVal := tmpSecret.getField(choice)
		if currentVal != "" {
			promptText += " [" + currentVal + "]"
		}

		value, _ = promptForText(promptText)
		value = strings.TrimSpace(value)
		if len(value) > 0 {
			field, _ := tmpSecret.assignValueToSecretStringField(choice, value)
			if field == choice {
				changed = true
			}
		}
	}

	// If nothing has changed, return
	if !changed {
		return
	}

	// I want to ask for confirmation here, so in case of deny, changes are not applied
	confirm, _ := promptForConfirm("Save")
	if !confirm {
		fmt.Println("Changes not saved")
		return
	}

	currentSecret.deepCopy(&tmpSecret)
	currentSecret.UpdatedAt = time.Now().Format(dateTimeFormat)

	saveVault(app, false)
	clearScreen()
}

func deleteSecretHandling(app *Context) {
	keys := make([]string, 0)
	for k := range app.vault.KeysMap {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		fmt.Println("No items in this vault")
		promptToJustWait()
		clearScreen()
		return
	}

	_, key, _ := promptForSelect("Choose", keys)

	currentSecret := app.vault.KeysMap[key]

	// I want to ask for confirmation here, so in case of deny, changes are not applied
	confirm, _ := promptForConfirm("Save")
	if !confirm {
		fmt.Println("Changes not saved")
		return
	}

	err := app.vault.delete(currentSecret.Name)
	if err != nil {
		check(err)
	}

	saveVault(app, false)
	clearScreen()
}

func getSecretHandling(app *Context) {
	keys := make([]string, 0)
	for k := range app.vault.KeysMap {
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
	if err != nil {
		return err
	}
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
		if !confirm {
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
