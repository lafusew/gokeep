/*
Copyright © 2022 LAFUSEW <antoine.oddoz@protonmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/lafusew/gokeep/data"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a credential record",
	Long:  `Delete a credential record`,
	Run: func(cmd *cobra.Command, args []string) {
		deletePrompt()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deletePrompt() {
	domainPromptContent := PromptContent{
		ErrorMsg: "This can't be empty",
		Label:    "Name of the creds you want to delete:",
	}

	var cred data.CredID

	err := TwoStepsSelect(domainPromptContent, &cred)

	if err != nil {
		fmt.Println(err)
		return
	}

	data.DeleteCred(cred)
}
