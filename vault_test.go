package main

import (
	"fmt"
	"testing"
	"time"
)

const (
	dateTimeFormat string = "2006-01-02T15:04:05-0700"
)

func TestVaultMarshaling(t *testing.T) {
	v := vault{}

	jsonString := `{
		"secrets": [
			{ "name": "GMAIL", "email": "test@email.com", "username": "alainrk", "password": "foobar", "apiKey": "3894H8ETW", "notes": "This is a secret", "createdAt": "2016-01-12T20:04:05-0700" },
			{ "name": "JIRA", "email": "test@email.com", "username": "jita", "password": "batterystaple", "apiKey": null, "notes": "Jira secret go on", "createdAt": "2020-02-19T15:12:05-0700" }
		]
	}`

	v.unmarshal(jsonString)
	// fmt.Println("Unmarshaled:", v)

	given := v.Secrets[0].Name
	expected := "GMAIL"
	if given != expected {
		t.Errorf("Wrong unmarshal. Given = %s, Expected = %s", given, expected)
	}

	given = v.Secrets[1].Email
	expected = "test@email.com"
	if given != expected {
		t.Errorf("Wrong unmarshal. Given = %s, Expected = %s", given, expected)
	}

	given = v.Secrets[1].ApiKey
	expected = ""
	if given != expected {
		t.Errorf("Wrong unmarshal. Given = %s, Expected = %s", given, expected)
	}

	marshaled := v.marshal()
	// fmt.Println("Marshaled:", marshaled)

	unmarshaledVault := vault{}
	unmarshaledVault.unmarshal(marshaled)

	given = v.Secrets[0].Password
	expected = unmarshaledVault.Secrets[0].Password
	if given != expected {
		t.Errorf("Marshal/Unmarshal test failed. Given = %s, Expected = %s", given, expected)
	}
}

func TestVaultOperations(t *testing.T) {
	v := vault{}
	jsonString := `{ "secrets": [] }`
	v.unmarshal(jsonString)

	given, _ := v.len()
	var expected int = 0
	if given != expected {
		t.Errorf("Failed vault length retrieval. Given = %d, Expected = %d", given, expected)
	}

	s := secret{}
	s.Name = "Lorem"
	s.Username = "ipsum"
	s.Email = "dolor@s.it"
	s.Password = "amet"
	s.ApiKey = "0398509234"
	s.Notes = "Test 1"
	s.CreatedAt = time.Now().Format(dateTimeFormat)
	v.add(s)

	s.Name = "This is another"
	s.Username = "Different item from the previous one"
	v.add(s)

	// fmt.Println("Insert secret:", s)
	fmt.Println(v)

	given, _ = v.len()
	expected = 2
	if given != expected {
		t.Errorf("Failed vault length retrieval after adding. Given = %d, Expected = %d", given, expected)
	}
}
