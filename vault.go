package main

import (
	"encoding/json"
	"errors"
)

const (
	notUnmarshaledError      string = "vault is not unmarshaled"
	secretAlreadyExistsError string = "secret name already exists in the vault"
)

type llNode struct {
	key  string
	next *llNode
}

type secret struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ApiKey    string `json:"apiKey"`
	Notes     string `json:"notes"`
	CreatedAt string `json:"createdAt"` // [GO= 2006-01-02T15:04:05-0700] [Java= yyyy-MM-dd'T'HH:mm:ssZ]	[C= %FT%T%z] ISO 8601
	UpdatedAt string `json:"updatedAt"` // [GO= 2006-01-02T15:04:05-0700] [Java= yyyy-MM-dd'T'HH:mm:ssZ]	[C= %FT%T%z] ISO 8601
}

type vault struct {
	name    string
	path    string
	KeysMap map[string]*secret // O(1) runtime mapping
	Secrets []secret           `json:"secrets"`
}

func (v *vault) unmarshal(jsonString string) {
	json.Unmarshal([]byte(jsonString), v)

	v.KeysMap = make(map[string]*secret)
	for _, secret := range v.Secrets {
		v.KeysMap[secret.Name] = &secret
	}
}

// Get the string out of the vault
func (v vault) marshal() string {
	x, _ := json.Marshal(v)
	return string(x)
}

func (v *vault) add(s secret) error {
	if _, ok := v.KeysMap[s.Name]; ok {
		return errors.New(secretAlreadyExistsError)
	}
	if v.Secrets == nil {
		v.Secrets = make([]secret, 1)
	}
	v.Secrets = append(v.Secrets, s)
	v.KeysMap[s.Name] = &s
	return nil
}

func (v vault) len() (int, error) {
	return len(v.Secrets), nil
}

// OLD Version, maybe useful in future
// Returns all the keys (names) of the vault, alphabetical sorted
func (v vault) getKeys() ([]string, error) {
	l, _ := v.len()
	keys := make([]string, l)

	head := llNode{"__head__", nil}
	for i := range v.Secrets {
		n := llNode{v.Secrets[i].Name, nil}

		if head.next == nil {
			head.next = &n
			continue
		}

		curr := &head
		// TODO: Binary search here
		for {
			if curr.next == nil {
				curr.next = &n
				break
			}
			if n.key <= curr.next.key {
				n.next = curr.next
				curr.next = &n
				break
			}
			curr = curr.next
		}
	}

	curr := head.next
	i := 0
	for curr != nil {
		keys[i] = curr.key
		curr = curr.next
		i += 1
	}

	return keys, nil
}

// Reflect alternative - Wrapper for field assignment
// Return fieldName if matched, error otherwise
func (s *secret) assignValueToSecretStringField(fieldName string, value string) (string, error) {
	switch fieldName {
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
	default:
		return "", errors.New("given field does not exist")
	}
	return fieldName, nil
}

// Constructors

func newVaultEmpty() *vault {
	v := vault{"", "", map[string]*secret{}, make([]secret, 0)}
	return &v
}
