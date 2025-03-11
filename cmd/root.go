/*
Copyright © 2024 nu12
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var tmpDir string
var rootCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Append PDF files",
	Long:  `Append PDF files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(appendCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&tmpDir, "temporary-directory", "t", "/tmp", "Temporary directory")
}
