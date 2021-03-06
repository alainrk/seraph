package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	password    string
	encryptMode bool
	decryptMode bool
}

// implement stringFlag to distinguish empty string given
type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

// Prompt the output on stdin and returns the clean input string given from the user
func prompt(output string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(output)
	text, err := reader.ReadString('\n')
	text = strings.TrimRight(text, "\n\r")
	if err != nil {
		panic(err)
	}
	return text
}

func getFlags() (flagsObject *flags) {
	var encryptFlag bool
	var decryptFlag bool
	var passwordFlag stringFlag

	flag.BoolVar(&encryptFlag, "e", false, "encrypt mode")
	flag.BoolVar(&decryptFlag, "d", false, "decrypt mode")
	flag.Var(&passwordFlag, "p", "password [INSECURE method, use interactive mode instead]")

	flag.Parse()

	if encryptFlag && decryptFlag {
		fmt.Println("Choose only one between [-d] and [-e] modality")
		os.Exit(1)
	}

	if !(encryptFlag || decryptFlag) {
		fmt.Println("Choose at least one between [-d] and [-e] modality")
		os.Exit(1)
	}

	flagsObject = &flags{}

	if encryptFlag {
		flagsObject.encryptMode = true
	}
	if decryptFlag {
		flagsObject.decryptMode = true
	}

	if !passwordFlag.set {
		// password not given
	} else if passwordFlag.value == "" {
		fmt.Println("Flag [-p] needs a string (password) argument")
		os.Exit(1)
	} else {
		flagsObject.password = passwordFlag.value
	}

	return flagsObject
}
