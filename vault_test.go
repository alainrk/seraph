package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVaultMarshaling(t *testing.T) {
	v := newVaultEmpty()

	jsonString := `{
		"secrets": [
			{ "name": "GMAIL", "email": "test@email.com", "username": "alainrk", "password": "foobar", "apiKey": "3894H8ETW", "notes": "This is a secret", "createdAt": "2016-01-12T20:04:05-0700", "updatedAt": "2020-02-19T15:12:05-0700" },
			{ "name": "JIRA", "email": "test@email.com", "username": "jita", "password": "batterystaple", "apiKey": null, "notes": "Jira secret go on", "createdAt": "2020-02-19T15:12:05-0700", "updatedAt": "2020-02-19T15:12:05-0700" }
		]
	}`

	v.unmarshal(jsonString)

	assert.Equal(t, v.Secrets[0].Name, "GMAIL", "they should be equal")

	assert.Equal(t, v.Secrets[1].Email, "test@email.com", "they should be equal")

	assert.Equal(t, v.Secrets[1].ApiKey, "", "they should be equal")

	marshaled := v.marshal()
	unmarshaledVault := newVaultEmpty()
	unmarshaledVault.unmarshal(marshaled)

	assert.Equal(t, v.Secrets[0].Password, unmarshaledVault.Secrets[0].Password, "they should be equal")
}

func TestVaultOperations(t *testing.T) {
	v := newVaultEmpty()
	jsonString := `{ "secrets": [] }`
	v.unmarshal(jsonString)

	len, _ := v.len()
	assert.Equal(t, len, 0, "they should be equal")

	s := secret{}
	s.Name = "Lorem"
	s.Username = "ipsum"
	s.Email = "dolor@s.it"
	s.Password = "amet"
	s.ApiKey = "0398509234"
	s.Notes = "Test 1"
	s.CreatedAt = time.Now().Format(dateTimeFormat)
	s.UpdatedAt = time.Now().Format(dateTimeFormat)
	v.add(s)

	s.Name = "This is another"
	s.Username = "Different item from the previous one"
	v.add(s)

	len, _ = v.len()
	assert.Equal(t, len, 2, "they should be equal")

	s.Name = "AKey number one"
	v.add(s)
	s.Name = "Das key nummer zwei"
	v.add(s)
	s.Name = "Ze kous nomiro dri"
	v.add(s)
	keys, err := v.getKeys()

	assert.Nil(t, err)

	assert.Equal(t, keys[2], "Lorem", "they should be equal")
}

func TestEmptyVault(t *testing.T) {
	v := newVaultEmpty()
	jsonString := `{ "secrets": [] }`
	v.unmarshal(jsonString)

	keys, _ := v.getKeys()
	assert.Equal(t, len(keys), 0, "they should be equal")
}

func TestSecretAssignment(t *testing.T) {

	s := secret{}
	s.Name = "Lorem"
	s.Email = "dolor@s.it"

	assert.Equal(t, s.Name, "Lorem", "they should be equal")

	s.assignValueToSecretStringField("Name", "Elon")
	s.assignValueToSecretStringField("Email", "elon@gmail.com")

	assert.Equal(t, s.Name, "Elon", "they should be equal")
	assert.Equal(t, s.Email, "elon@gmail.com", "they should be equal")
}
