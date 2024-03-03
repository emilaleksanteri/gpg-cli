/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gpg-gen/util"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List GPG keys",
	Long:  `List long format of all generated GPG keys on this machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := util.GetCurrentDir()
		if err != nil {
			fmt.Printf("Error getting current dir: %v\n", err)
			return
		}

		err = util.ExecuteCmdWithPrint(
			"gpg",
			[]string{"--list-secret-keys", "--keyid-format=long"},
			dir,
		)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}
