/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gpg-gen/util"
)

func init() {
	var key string
	var email string

	addCmd.Flags().StringVar(&key, "key", "", "key id for a GPG key, required")
	addCmd.Flags().StringVar(&email, "email", "", "email for a GPG key, required")
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add GPG key config to your local git config, flags --key and --email are required",
	Long:  `Adds previously generated GPG key to local git config and changes local email for the git to email associated with the key. Needs --key and --email flags`,
	Run: func(cmd *cobra.Command, args []string) {
		keyFlag := cmd.Flag("key").Value.String()
		if keyFlag == "" {
			fmt.Println("Key flag is required")
			return
		}

		emailFlag := cmd.Flag("email").Value.String()
		if emailFlag == "" {
			fmt.Println("Email flag is required")
			return
		}

		currDir, err := util.GetCurrentDir()
		if err != nil {
			fmt.Printf("Error getting current dir: %v\n", err)
			return
		}

		if !util.CheckGitExistsOnDir(currDir) {
			fmt.Printf("Git repo not found on dir, run: git init\n")
			return
		}

		err = util.ExecuteCmd(
			"git",
			[]string{"config", "--unset", "gpg.format"},
			currDir,
		)

		err = util.ExecuteCmd(
			"git",
			[]string{"config", "user.signingkey", keyFlag},
			currDir,
		)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		err = util.ExecuteCmd(
			"git",
			[]string{"config", "commit.gpgsign", "true"},
			currDir,
		)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

		err = util.ExecuteCmd(
			"git",
			[]string{"config", "user.email", emailFlag},
			currDir,
		)

		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
