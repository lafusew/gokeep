/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/lafusew/gokeep/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new credentials database and table",
	Long:  `Init a new credentials database and table. WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		initialize()
	},
}

func initialize() {
	err := data.CreateCredsTable()
	if err != nil {
		fmt.Println(err.Error())
	}

	// pc := PromptContent{
	// 	Label:    "Set up a phrase that you can read. You'll later use this to be able to confirm that u entered the correct key",
	// 	ErrorMsg: "This can't be empty (should be readable)",
	// }

	// key, err := PromptGetInput(pc)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// bkey := []byte(key)
	// f, err := os.Create("/data/key")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer f.Close()

	// err = os.WriteFile("/data/key", bkey, fs.FileMode(0644))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
