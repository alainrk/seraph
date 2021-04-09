package main

import (
	"encoding/json"
)

type secret struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ApiKey    string `json:"apiKey"`
	Notes     string `json:"notes"`
	CreatedAt string `json:"createdAt"`
}

type vault struct {
	Name    string
	Email   string
	Secrets []secret `json:"secrets"`
}

func (v *vault) unmarshal(jsonString string) {
	json.Unmarshal([]byte(jsonString), v)
}

func (v *vault) marshal() string {
	x, _ := json.Marshal(v)
	return string(x)
}
