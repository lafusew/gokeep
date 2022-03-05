/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lafusew/gokeep/data"
	"github.com/spf13/cobra"
)

// createCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new credential record",
	Long: `Create s new credential record. 
	
	By default you'll need to provide domain, username and password.
	If you want gokeep to generates a password for you, tbw
	`,
	Run: func(cmd *cobra.Command, args []string) {
		createPrompt()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func createPrompt() {
	domainPromptContent := PromptContent{
		ErrorMsg: "This can't be empty, please provide a domain name",
		Label:    "Service's name for which you want to store credentials:",
	}

	domain, err := PromptGetInput(domainPromptContent)

	if err != nil {
		return
	}

	usernamePromptContent := PromptContent{
		ErrorMsg: "This can't be empty, please provide a identifier, it can be anything (mail, phone numbre, username)",
		Label:    "Credentials identifier you use to log in:",
	}

	username, err := PromptGetInput(usernamePromptContent)

	if err != nil {
		return
	}

	pwdPromptContent := PromptContent{
		ErrorMsg: "This can't be empty, please provide a password",
		Label:    "Password used with this identifier:",
	}

	pwd, err := PromptGetInput(pwdPromptContent)

	if err != nil {
		return
	}

	data.InsertCred(domain, username, pwd)
}
