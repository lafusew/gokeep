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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	mk := data.GetMK()

	if mk == "" {
		setKeyPrompt()
	}

	var (
		l = "List"
		c = "Create"
		r = "Read"
		u = "Update"
		d = "Delete"
	)

	items := []string{l, c, r, u, d}

	prompt := promptui.Select{
		Label: "Manage your credentials",
		Items: items,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		log.Fatalf("Command cancelled \n%v\n", err)
	}

	switch selected {
	case l:
		readAllPrompt()
	case c:
		createPrompt()
	case r:
		readPrompt()
	case u:
		updatePrompt()
	case d:
		deletePrompt()
	default:
		readAllPrompt()
	}

	run()
}
