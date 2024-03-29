/*
Copyright © 2024 nu12
*/
package cmd

import (
	"os"

	"github.com/nu12/pdf/cmd/create"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdf",
	Short: "Create, append or split PDF files",
	Long:  `Create, append or split PDF files`,
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
	rootCmd.AddCommand(create.Cmd)
}
