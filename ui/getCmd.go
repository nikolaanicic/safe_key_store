package ui

import (
	"fmt"
	"safe_key_store/passmanager"

	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command{
	cmd := &cobra.Command{
		Use: "get",
		Short: "used to retrieve saved passwords",
		Long: "this command will retrieve passwords from windows credential manager",
		Run:get,
	}

	insertFlag(cmd,keyflagname)

	return cmd
}


func get(c *cobra.Command, args []string){

	passfv, err := getFlagValueFromCmd(c,keyflagname)
	if err != nil{
		fmt.Println(err)
		return
	}

	pass, err := passmanager.GetPassword(passfv)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(pass)

}