/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/lafusew/gokeep/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage master password",
	Long: `Manage master password`,
	Run: func(cmd *cobra.Command, args []string) {
		manageKey()
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func manageKey() {
	items := []string{
		"Set",
		"Forget",
		"Init",
		"Timeout",
	}

	prompt := promptui.Select{
		Label: "Manage your key",
		Items: items,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Fatalf("Command cancelled \n%v\n", err)
	}

	switch selected {
	case "Set":
		setKeyPrompt()
	case "Forget":
		data.SetMK("")
	}
}

func setKeyPrompt() {
	keyPromtContent := PromptContent{
		Label: "Your key (1-32 characters): ",
		ErrorMsg: "Can't be empty",
	}

	key, err := PromptGetInput(keyPromtContent)
	if err != nil {
		return
	}

	data.SetMK(key)
}
