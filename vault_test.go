package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testVaultPath = "./vaults/__onlytest.vault"

func truncateTestVault() {
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
	truncateTestVault()

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
