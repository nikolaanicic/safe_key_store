package ui

import "fmt"

func masterpassfailcheck() {

	loop := true
	
	for loop {
		masterpass, err := askMasterPassword()
		if err != nil{
			fmt.Println("invalid master password, please try again")
		} else {
			loop = !CheckMasterPassword(masterpass)
		}
	}
}