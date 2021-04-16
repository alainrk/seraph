package main

import (
	"github.com/manifoldco/promptui"
)

func promptForText(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func promptForTextValid(label string, validate func(string) error) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	check(err)

	return result
}

func promptForPassword(label string, validate func(string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	return result, err
}

func promptForSelect(label string, choices []string) (int, string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: choices,
	}

	index, mode, err := prompt.Run()

	if err != nil {
		return index, mode, err
	}

	return index, mode, nil
}

func promptForConfirm(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return false, nil
	}

	return result == "y", nil
}
