package main

import (
	"fmt"
)

const (
	dateTimeFormat string = "2006-01-02T15:04:05-0700"
)

const (
	vaultDirectory          = "./SeraphVaults/"
	vaultAlreadyExistsError = "vault already exists"
)

// Step 0
const (
	exit int = iota
	openVault
	newVault
	testVault
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Context struct {
	hashedPassword string
	vault          *vault
}

func main() {
	var app *Context
	clearScreen()
	initVaultDirectory()

	// TODO: Non-interactive handling
	// flags := getFlags()

	for {
		index, _, _ := promptForSelect("Choose", []string{"Exit", "Open Vault", "New Vault"})

		if index == exit {
			return
		}

		// Re-init at every cycle
		app = &Context{}

		if index == openVault {
			// Opening existing vault
			err := chooseVault(app)
			if err != nil {
				continue
			}
			openedVaultHandling(app)
		} else if index == newVault {
			// Create new vault
			err := newVaultHandling(app)
			if err != nil {
				fmt.Println(err)
				return
			}
			openedVaultHandling(app)
		}
	}
}
