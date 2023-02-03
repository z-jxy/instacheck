/*
Copyright © 2023 - zjxy
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var filePath string

var rootCmd = &cobra.Command{
	Use:   "instacheck",
	Short: "Automated verification of available Instagram usernames",
	Long: `Check if any usernames in a file are available or taken on instagram.
Usernames in the file should be formatted as just the username without any '@' prefix.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[*] Starting...")
		run(filePath)
		fmt.Println("[*] Finished!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to file containing usernames")
	rootCmd.MarkFlagRequired("file")
}
