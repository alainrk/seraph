package main

import (
	"encoding/json"
	"errors"
)

const (
	notUnmarshaledError string = "Vault is not unmarshaled"
)

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
