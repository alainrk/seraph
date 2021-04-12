package main

import (
	"encoding/json"
	"errors"
)

const (
	notUnmarshaledError string = "vault is not unmarshaled"
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
}

type vault struct {
	__isUnmarshaled bool
	Secrets         []secret `json:"secrets"`
}

func (v *vault) unmarshal(jsonString string) {
	json.Unmarshal([]byte(jsonString), v)
	v.__isUnmarshaled = true
}

func (v vault) marshal() string {
	x, _ := json.Marshal(v)
	return string(x)
}

func (v *vault) add(s secret) error {
	if !v.__isUnmarshaled {
		return errors.New(notUnmarshaledError)
	}
	v.Secrets = append(v.Secrets, s)
	return nil
}

func (v vault) len() (int, error) {
	if !v.__isUnmarshaled {
		return -1, errors.New(notUnmarshaledError)
	}
	return len(v.Secrets), nil
}

// Returns all the keys (names) of the vault, alphabetical sorted
func (v vault) getKeys() ([]string, error) {
	l, err := v.len()
	if err != nil {
		return make([]string, 0), errors.New(notUnmarshaledError)
	}
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
