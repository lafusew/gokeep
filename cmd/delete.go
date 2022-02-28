/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/lafusew/gokeep/data"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete a credential record",
	Long: `Delete a credential record`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteCred()
	},
}

func init() {
	credsCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteCred() {
	domainPromptContent := PromptContent{
		"This can't be empty",
		"Name of the creds you want to delete:",
	}

	domain, err := PromptGetInput(domainPromptContent)

	if err != nil {
		return
	}

	pDomains := data.FindCred(domain)

	if len(pDomains) < 1 {
		fmt.Println("No credentials found, command cancelled")
		return
	}

	res, err := PromptGetSelect(pDomains, "Confirm selection:")

	if err != nil {
		return
	}

	data.DeleteCred(res)
}


