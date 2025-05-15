/*
Copyright © 2024 nu12
*/

package cmd

import (
	pdf "github.com/nu12/pdf/pkg"
	"github.com/spf13/cobra"
)

func init() {}

var appendCmd = &cobra.Command{
	Use:   "append input1 [input2 input3 ... inputN] output",
	Short: "Append PDFs and images into a single PDF file",
	Long: `Append PDFs and images into a single PDF file
	
Examples:

# Append two PDFs into a single PDF file:
pdf append input1.pdf input2.pdf output.pdf

# Append PDFs and images into a single PDF file:
pdf append input1.pdf image1.jpg output.pdf

# Append all files from a folder into a single PDF file:
pdf append path-to-files/* output.pdf`,

	Run: func(cmd *cobra.Command, args []string) {
		pdf.InitAppend(args, tmpDir).PreProcess().Run().PostProcess()
	},
}
