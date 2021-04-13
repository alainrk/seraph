package main

import "errors"

const (
	tooShortPasswordError string = "passphrase is too short"
)

func validatePassword(s string) error {
	if len(s) < 8 {
		return errors.New(tooShortPasswordError)
	}
	return nil
}

func validateAlwaysString(s string) error {
	return nil
}
