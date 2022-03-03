/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/lafusew/gokeep/data"
	"github.com/lafusew/gokeep/prompter"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a credential record",
	Long:  `Read a credential record`,
	Run: func(cmd *cobra.Command, args []string) {
		readCred()
	},
}

func readCred() error {
	domainPromptContent := prompter.PromptContent{
		ErrorMsg: "This can't be empty, please provide a domain name",
		Label:    "Service's name for which you want to retrieve credentials:",
	}
	var cred data.CredID

	err := prompter.TwoStepsSelect(domainPromptContent, &cred)

	res := data.FindCredById(cred.Id)

	fmt.Println(res)

	return err
}

func init() {
	credsCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle
}
