/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gpg-gen/util"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Generate a copyable GPG key format for GitHub for a GPG key ID",
	Long: `
	Takes GPG key ID as a flag and generates a copyable GPG key format for GitHub.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		key := cmd.Flag("key").Value.String()
		if key == "" {
			fmt.Println("key is required")
			return
		}
		currDir, err := util.GetCurrentDir()
		if err != nil {
			fmt.Printf("Error getting current directory: %v", err)
			return
		}

		fmt.Printf("\n\nCOPY ALL OF THE FOLLOWING: \n\n")
		err = util.ExecuteCmdWithPrint(
			"gpg",
			[]string{"--armor", "--export", key},
			currDir,
		)

		if err != nil {
			fmt.Printf("Error exporting key: %v", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	var key string

	githubCmd.Flags().StringVar(&key, "key", "", "key id for a GPG key, required")
}
