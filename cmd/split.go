/*
Copyright © 2024 nu12
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Dir string

func init() {
	splitCmd.Flags().StringVarP(&Dir, "dir", "d", "./", "Output directory.")
}

var splitCmd = &cobra.Command{
	Use:   "split input.pdf -d /path/to/output",
	Short: "Split a PDF file into multiple PDF files",
	Long:  `Split a PDF file into multiple PDF files.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Split called")

	},
}
