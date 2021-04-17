package main

import "fmt"

const (
	dateTimeFormat string = "2006-01-02T15:04:05-0700"
)

const (
	vaultDirectory          = "./vaults/"
	vaultAlreadyExistsError = "vault already exists"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Context struct {
	hashedPassword string
	vault          *vault
	currentStep    string
}

func main() {
	var ctx *Context

	// TODO: Non-interactive handling
	// flags := getFlags()

	// Step 0
	const (
		exit int = iota
		openVault
		newVault
		testVault
	)

	for {
		index, _, _ := promptForSelect("Choose", []string{"Exit", "Open Vault", "New Vault"})

		if index == exit {
			return
		}

		// Re-init at every cycle
		ctx = &Context{}

		if index == openVault {
			// Opening existing vault
			err := chooseVault(ctx)
			if err != nil {
				return
			}
			openedVaultHandling(ctx)
		} else if index == newVault {
			// Create new vault
			err := newVaultHandling(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			openedVaultHandling(ctx)
		}
	}
}
