package ui

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd cobra.Command

// initializes the root cobra cli command
// with subcommands
func initRootCmd(commands... *cobra.Command) {
	rootCmd = cobra.Command{
		Use: 	"passmgr",
		Short:	"pass is a simple password manager program for windows",
		Long: 	"pass is a password manager that uses some windows based functionalities",

	}

	rootCmd.AddCommand(commands...)

}



func Execute() {
	if err := rootCmd.Execute(); err != nil{
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func InitUI(){
	initRootCmd(NewStoreCredCmd(),NewGetCmd())
}