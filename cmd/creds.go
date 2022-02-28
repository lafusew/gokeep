/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// credsCmd represents the creds command
var credsCmd = &cobra.Command{
	Use:   "creds",
	Short: "Credentials are stored and accessed through the gokeep creds command",
	Long: `Credentials are stored and accessed through the gokeep creds command`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	domainName, _ := cmd.Flags().GetString("d")
	// 	log.Print("just ran creds cmd")

	// 	log.Print(domainName)
	// },
}

func init() {
	rootCmd.AddCommand(credsCmd)
}