/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// credsCmd represents the creds command
var credsCmd = &cobra.Command{
	Use:   "creds",
	Short: "Credentials are stored and accessed through the gokeep creds command",
	Long:  `Credentials are stored and accessed through the gokeep creds command`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	items := []string{
		"List",
		"Create",
		"Read",
		"Update",
		"Delete",
	}

	prompt := promptui.Select{
		Label: "Manage your credentials",
		Items: items,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Fatalf("Command cancelled \n%v\n", err)
	}

	switch selected {
	case "Create":
		createNewCred()
	case "Delete":
		deleteCred()
	case "Read":
		readCred()
	}

	run()
}

func init() {
	rootCmd.AddCommand(credsCmd)
}
