/*
Copyright © 2024 nu12
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `Show current version`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
