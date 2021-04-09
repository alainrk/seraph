package main

import (
	"fmt"
	"testing"
)

func TestVault(t *testing.T) {
	v := vault{}

	jsonString := `{ "secrets": [{ "name": "GMAIL", "email": "test@email.com", "username": "alainrk", "password": "foobar", "apiKey": "3894H8ETW", "notes": "This is a secret", "createdAt": "2020-02-12 09:12:32" },{ "name": "JIRA", "email": "test@email.com", "username": "jita", "password": "batterystaple", "apiKey": null, "notes": "Jira secret go on", "createdAt": "2020-02-14 09:12:32" }] }`

	v.unmarshal(jsonString)

	fmt.Println("Unmarshaled:", v)

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
	fmt.Println("Marshaled:", marshaled)

	unmarshaledVault := vault{}
	unmarshaledVault.unmarshal(marshaled)

	given = v.Secrets[0].Password
	expected = unmarshaledVault.Secrets[0].Password
	if given != expected {
		t.Errorf("Marshal/Unmarshal test failed. Given = %s, Expected = %s", given, expected)
	}
}
