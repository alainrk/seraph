package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testVaultPath = "./vaults/__onlytest.vault"

func truncateVault() {
	f, err := os.OpenFile(testVaultPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestVaultCreateStub(t *testing.T) {
	v := newVaultEmpty()
	truncateVault()

	jsonString := `{
		"secrets": [
			{ "name": "GMAIL", "email": "gmail@email.com", "username": "alainrk", "password": "foobar", "apiKey": "3894H8ETW", "notes": "This is a secret", "createdAt": "2016-01-12T20:04:05-0700", "updatedAt": "2020-02-19T15:12:05-0700" },
			{ "name": "JIRA", "email": "jira@email.com", "username": "jitark", "password": "batterystaple", "apiKey": null, "notes": "Jira secret go on", "createdAt": "2020-02-19T15:12:05-0700", "updatedAt": "2020-02-19T15:12:05-0700" },
			{ "name": "GITHUB", "email": "github@email.com", "username": "gitbabbo", "password": "12345", "apiKey": null, "notes": "Github secret go on", "createdAt": "2020-02-19T15:12:05-0700", "updatedAt": "2020-02-19T15:12:05-0700" }
		]
	}`

	err := v.unmarshal(jsonString)
	assert.Nil(t, err, "there should not be an error in unmarshalling")

	assert.Equal(t, v.KeysMap["GMAIL"].Name, "GMAIL", "they should be equal")
	assert.Equal(t, v.KeysMap["GMAIL"].Email, "gmail@email.com", "they should be equal")
	assert.Equal(t, v.KeysMap["JIRA"].Username, "jitark", "they should be equal")
	assert.Equal(t, v.KeysMap["GITHUB"].Username, "gitbabbo", "they should be equal")

	marshaledPlainText := v.marshal()
	hashedPassword := hashPassword("password")
	marshaledCipherText := encrypt(hashedPassword, marshaledPlainText)
	err = ioutil.WriteFile(testVaultPath, []byte(marshaledCipherText), 0644)

	assert.Nil(t, err, "they should be equal")
}

func TestSecretDeepCopy(t *testing.T) {
	fakeVal := "xxx"
	s := secret{}
	s.Name = fakeVal
	s.Username = fakeVal
	s.Email = fakeVal
	s.Password = fakeVal
	s.ApiKey = fakeVal
	s.Notes = fakeVal
	s.CreatedAt = fakeVal
	s.UpdatedAt = fakeVal

	d := secret{}
	d.deepCopy(&s)

	assert.Equal(t, d.Name, s.Name, "they should be equal")
	assert.Equal(t, d.Username, s.Username, "they should be equal")
	assert.Equal(t, d.Email, s.Email, "they should be equal")
	assert.Equal(t, d.Password, s.Password, "they should be equal")
	assert.Equal(t, d.ApiKey, s.ApiKey, "they should be equal")
	assert.Equal(t, d.Notes, s.Notes, "they should be equal")
	assert.Equal(t, d.CreatedAt, s.CreatedAt, "they should be equal")
	assert.Equal(t, d.UpdatedAt, s.UpdatedAt, "they should be equal")
}

func TestReflectVault(t *testing.T) {
	fakeVal := "xxx"
	s := secret{}
	s.Name = fakeVal
	s.Username = fakeVal
	s.Email = fakeVal
	s.Password = fakeVal
	s.ApiKey = fakeVal
	s.Notes = fakeVal
	s.CreatedAt = fakeVal
	s.UpdatedAt = fakeVal

	assert.Equal(t, s.getField("Name"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("Username"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("Email"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("Password"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("ApiKey"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("Notes"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("CreatedAt"), fakeVal, "they should be equal")
	assert.Equal(t, s.getField("UpdatedAt"), fakeVal, "they should be equal")
}
