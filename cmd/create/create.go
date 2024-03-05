/*
Copyright © 2024 nu12
*/

package create

import (
	"github.com/spf13/cobra"
)

var Margin string
var Compression string
var Orientation string

func init() {
	Cmd.Flags().StringVarP(&Compression, "compression", "c", "NONE", "Define the amount of compression to be applied to the input images. Available values are: NONE, LOWEST, LOW, MEDIUM, HIGH, HIGHEST.")
	Cmd.Flags().StringVarP(&Margin, "margin", "m", "NONE", "Define the margins for the output document. Available values are: NONE, MIN, MEDIUM, MAX.")
	Cmd.Flags().StringVarP(&Orientation, "orientation", "o", "P", "Define the orientation of the output document. Available values are: P, L.")
}

var Cmd = &cobra.Command{
	Use:   "create input1 [input2 input3 ... inputN] output",
	Short: "Create a PDF file from images",
	Long:  `Create a PDF file from images (png).`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}
