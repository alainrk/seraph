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

	// Step 1
	const (
		back int = iota
		getSecret
		insertSecret
	)

	for {
		index, _, _ := promptForSelect("Choose", []string{"Exit", "Open Vault", "New Vault", "TEST-PrintVault"})

		if index == exit {
			return
		}

		// Re-init at every cycle
		ctx = &Context{}

		// TEST - Removeme
		if index == testVault {
			chooseVault(ctx)
			fmt.Println(ctx.vault.marshal())
		}

		if index == openVault {
			// Opening existing vault
			chooseVault(ctx)

			index, _, _ = promptForSelect("Choose", []string{"Back", "Get secret", "Insert secret"})

			if index == insertSecret {
				insertSecretHandling(ctx)
			} else if index == getSecret {
				getSecretHandling(ctx)
			}
		} else if index == newVault {
			// Create new vault
			newVaultHandling(ctx)
		}
	}
}
