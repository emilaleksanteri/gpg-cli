/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpg-gen",
	Short: "ClI app to add GPG keys to git config and github fast",
	Long: `
	Really this cli tool is just for doing the right steps to add GPG keys to your github account

	Before you start, to generate a GPG key with gpg command line run:
	$ gpg --full-generate-key
	
	Then you can run:
	$ gpg-gen keys

	To pick key id and email from the key list
	
	%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
	sec   4096R/3AA5C34371567BD2 2016-03-10 [expires: 2017-03-10]
	uid                          Hubot <hubot@example.com>       
	ssb   4096R/4BB6D45482678BE3 2016-03-10		             
	%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

	Key id would be: 3AA5C34371567BD2
	Email would be: hubot@example.com

	Then you can run:
	$ gpg-gen add --key <key-id> --email <email>

	To add the key to your git config and github account:
	$ gpg-gen github --key <key-id>

	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gpg-gen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
