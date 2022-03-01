/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lafusew/gokeep/data"
	"github.com/lafusew/gokeep/prompter"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new credential record",
	Long: `Create s new credential record. 
	
	By default you'll need to provide domain, username and password.
	If you want gokeep to generates a password for you, tbw
	`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewCred()
	},
}

func init() {
	credsCmd.AddCommand(newCmd)
}

func createNewCred() {
	domainPromptContent := prompter.PromptContent{
		ErrorMsg: "This can't be empty, please provide a domain name",
		Label: "Service's name for which you want to store credentials:",
	}

	domain, err := prompter.PromptGetInput(domainPromptContent)

	if err != nil {
		return
	}

	usernamePromptContent := prompter.PromptContent{
		ErrorMsg: "This can't be empty, please provide a identifier, it can be anything (mail, phone numbre, username)",
		Label: "Credentials identifier you use to log in:",
	}

	username, err := prompter.PromptGetInput(usernamePromptContent)

	if err != nil {
		return
	}

	pwdPromptContent := prompter.PromptContent{
		ErrorMsg: "This can't be empty, please provide a password",
		Label: "Password used with this identifier:",
	}

	pwd, err := prompter.PromptGetInput(pwdPromptContent)

	if err != nil {
		return
	}

	data.InsertCred(domain, username, pwd)
}