/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/lafusew/gokeep/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a credentials record",
	Long: `Update a credentials record`,
	Run: func(cmd *cobra.Command, args []string) {
		updatePrompt()
	},
}

func updatePrompt() {
	var cred data.CredID

	err := TwoStepsSelect(
		PromptContent{
			Label: "Credentials to update: ",
			ErrorMsg: "Can't be empty",
		},
		&cred,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	prompt := promptui.Select{
		Label: "field to update: ",
		Items: []string{"domain", "username", "password"},
	}

	_, field, err := prompt.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}

	label := fmt.Sprintf("New value for %s: ", field)
	value, err := PromptGetInput(PromptContent{ Label: label, ErrorMsg: "This can't be empty"})

	if err != nil {
		log.Fatalln(err.Error())
	}

	data.UpdateCred(cred.Id, field, value)

	prompt = promptui.Select{
		Label: "Continue to update this creds?",
		Items: []string{"yes", "no"},
	}

	_, continuation, err := prompt.Run()

	if err != nil {
		log.Fatalln(err.Error())
	}

	if continuation == "yes" {
		updatePrompt()
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
