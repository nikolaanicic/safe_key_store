package ui

import (
	"fmt"
	"safe_key_store/passmanager"
	"strings"
)

// asks the user to enter the master password
func askMasterPassword() string {
	pass := ""

	for pass == ""{
		fmt.Println("Please enter the admin password:")
		fmt.Scanf("%s",&pass)
		if err := validateInput(pass); err != nil{
			fmt.Println(err)
			pass = ""
		}
	}

	return pass
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


