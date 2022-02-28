/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/lafusew/gokeep/data"
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

type PromptContent struct {
	errorMsg string
	label string
}

func init() {
	credsCmd.AddCommand(newCmd)
}

func createNewCred() {
	domainPromptContent := PromptContent{
		"This can't be empty, please provide a domain name",
		"Service's name for which you want to store credentials:",
	}

	domain, err := PromptGetInput(domainPromptContent)

	if err != nil {
		return
	}

	usernamePromptContent := PromptContent{
		"This can't be empty, please provide a identifier, it can be anything (mail, phone numbre, username)",
		"Credentials identifier you use to log in:",
	}

	username, err := PromptGetInput(usernamePromptContent)

	if err != nil {
		return
	}

	pwdPromptContent := PromptContent{
		"This can't be empty, please provide a password",
		"Password used with this identifier:",
	}

	pwd, err := PromptGetInput(pwdPromptContent)

	if err != nil {
		return
	}

	data.InsertCred(domain, username, pwd)
}