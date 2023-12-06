package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "gowscat",
	Short:   "set request Method",
	Long:    "set a request method is support post、get",
	Example: "-M GET|POST",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre run method command")
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
