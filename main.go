package main

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
	ctx := Context{}
	// TODO: Non-interactive handling
	// flags := getFlags()

	// Step 0
	const (
		openVault int = iota
		newVault
	)

	// Step 1
	const (
		insertSecret int = iota
		getSecret
	)

	index, _, _ := promptForSelect("Choose", []string{"Open Vault", "New Vault"})

	// Opening vault
	if index == openVault {
		chooseVault(ctx)
		index, _, _ = promptForSelect("Choose", []string{"Get secret", "Insert secret"})
		if index == insertSecret {
			getSecretHandling(ctx)
		} else if index == getSecret {
			insertSecretHandling(ctx)
		}
	} else if index == newVault {
		newVaultHandling(ctx)
	}
}
