/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/lafusew/gokeep/data"
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
	domainPromptContent := PromptContent{
		ErrorMsg: "This can't be empty, please provide a domain name",
		Label:    "Service's name for which you want to retrieve credentials:",
	}
	var cred data.CredID

	err := TwoStepsSelect(domainPromptContent, &cred)

	res := data.FindCredById(cred.Id)

	fmt.Println(res)

	return err
}

func readAllCreds() error {
	creds := data.FindAllCreds()

	cred, err := PromptGetSelect(creds, "Select a credential:")

	if err != nil {
		return err
	}

	res := data.FindCredById(cred.Id)

	fmt.Println(res)

	return err
}

func init() {
	credsCmd.AddCommand(readCmd)
}
