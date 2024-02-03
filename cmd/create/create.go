/*
Copyright © 2024 nu12
*/

package create

import (
	"fmt"
	"os"
	"strings"

	"github.com/nu12/pdf/internal/helper"
	"github.com/nu12/pdf/internal/model"
	pdf "github.com/nu12/pdf/pkg"
	"github.com/spf13/cobra"
)

var availableCompressionProfiles = map[string]model.CompressionProfile{
	"NONE":    {Value: 1},
	"LOWEST":  {Value: 2},
	"LOW":     {Value: 3},
	"MEDIUM":  {Value: 4},
	"HIGH":    {Value: 5},
	"HIGHEST": {Value: 6},
}

var availableMarginProfiles = map[string]model.MarginProfile{
	"NONE":   {X: 0, Y: 0},
	"MIN":    {X: 10, Y: 10},
	"MEDIUM": {X: 20, Y: 20},
	"MAX":    {X: 40, Y: 40},
}

var availableModeProfiles = map[string]model.ModeProfile{
	"P": {Format: "A4", Orientation: "P", Width: 210, Height: 297},
	"L": {Format: "A4", Orientation: "L", Width: 297, Height: 210},
}

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
	Short: "Create PDF files from images",
	Long:  `Create PDF files from images (png).`,

	Run: func(cmd *cobra.Command, args []string) {

		length := len(args)

		if length < 2 {
			fmt.Println("Not enough arguments.")
			os.Exit(1)
		}

		upperLimit := length - 1
		inputs := args[:upperLimit]
		output := args[upperLimit]

		selectedMarginProfile := availableMarginProfiles[strings.ToUpper(Margin)]
		selectedCompressionProfile := availableCompressionProfiles[strings.ToUpper(Compression)]
		selectedModeProfile := availableModeProfiles[strings.ToUpper(Orientation)]

		files := helper.UnmarshalImages(inputs)

		if err := pdf.Create(files, output, selectedMarginProfile, selectedCompressionProfile, selectedModeProfile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
