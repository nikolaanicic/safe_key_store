package ui

import (
	"github.com/spf13/cobra"
)

const(
	keyflagname = "key"
	passwordflagname = "pass"
	helpflagname = "help"
)


// gets the value from the command
func getFlagValueFromCmd(cmd *cobra.Command, name string) (value string, err error){
	value, err = cmd.Flags().GetString(name)
	if err == nil{
		err = validateInput(value)
	}
	return
}

func insertFlag(cmd *cobra.Command, flagname string){
	switch flagname{
	case keyflagname:
		cmd.Flags().String(keyflagname,"","pass the key/username for a credential")
	case passwordflagname:
		cmd.Flags().String(passwordflagname,"","pass the password for a credential")
	}
}
