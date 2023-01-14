package ui

import (
	"fmt"
	"safe_key_store/passmanager"
	"strings"
	"syscall"

	"golang.org/x/term"
)

// asks the user to enter the master password
func askMasterPassword() (string,error) {
	fmt.Println("Please enter the master password:")
	pass, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil{
		return "", err
	}
	return string(pass),nil
}

// checks the master passwrod entered by the user
func CheckMasterPassword(pass string) bool {
	if err := passmanager.VerifyMasterPassword(pass); err != nil{
		fmt.Println(err.Error())
		return false
	}
	return true
}

func validateInput(input string) error {
	input = strings.TrimSpace(input)
	if input == ""{
		return fmt.Errorf("empty")
	} else if len(input) < 2 {
		return fmt.Errorf("shorter than two characters")
	}

	return nil

}


// func validateKey(key string) error{
// 	return validateInput(key)
// }


