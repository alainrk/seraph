package main

import (
	"errors"
)

const (
	tooShortPasswordError     string = "password is too short"
	passwordsNotMatchingError string = "passwords do not match"
)

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New(tooShortPasswordError)
	}
	return nil
}
