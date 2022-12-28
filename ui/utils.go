package ui

func masterpassfailcheck() {
	if !CheckMasterPassword(askMasterPassword()) {
		panic("invalid master password")
	}
}