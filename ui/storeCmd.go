package ui

import (
	"fmt"
	"safe_key_store/passmanager"

	"github.com/spf13/cobra"
)

func NewStoreCredCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: 	"store [FLAGS] (--key | --pass)",
		Short: 	"store - stores a credential in the windows credential manager store",
		Long:	"store - used for storing new credentials",
		Run: store,
	}

	insertFlag(cmd,keyflagname)
	insertFlag(cmd,passwordflagname)

	return cmd
}


func store(cmd *cobra.Command, args []string){

	masterpassfailcheck()
	keyval, err := getFlagValueFromCmd(cmd,keyflagname)
	if err != nil{
		fmt.Println(keyflagname,err)
		return
	}

	passval, err := getFlagValueFromCmd(cmd,passwordflagname)
	if err != nil{
		fmt.Println(passwordflagname,err)
		return
	}

	if err := passmanager.StoreNewCredential(keyval,passval); err != nil{
		fmt.Println(err)
		return
	}
}