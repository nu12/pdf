/*
Copyright © 2024 nu12
*/

package append

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var Compression string

func init() {
	Cmd.Flags().StringVarP(&Compression, "compression", "c", "NONE", "Define the amount of compression to be applied to the input images. Available values are: NONE, LOWEST, LOW, MEDIUM, HIGH, HIGHEST.")
}

var Cmd = &cobra.Command{
	Use:   "append input1 [input2 input3 ... inputN] output",
	Short: "Append PDFs and images into a single PDF file",
	Long:  `Append PDFs and images into a single PDF file.`,

	Run: func(cmd *cobra.Command, args []string) {

		if _, err := exec.LookPath("gs"); err != nil {
			fmt.Println("gs executable not found")
			return
		}

		if len(args) < 2 {
			fmt.Println("You must provide at least two arguments: input file or files and an output file")
			return
		}

		output := args[len(args)-1]
		inputs := args[:len(args)-1]

		// TODO: Add resolution setting (-r150)
		cmdArgs := []string{
			"gs",
			"-q",
			"-dNOPAUSE",
			"-dBATCH",
			"-sDEVICE=pdfwrite",
			"-dCompatibilityLevel=1.4",
			"-dPDFSETTINGS=/default",
			"-dCompressFonts=true",
			"-sOutputFile=" + output,
			strings.Join(inputs, " "),
		}

		// TODO: Fix error when file name contains spaces
		// TODO: Convert images to PDFs before appending
		shellCmd := exec.Command("/bin/bash", "-c", strings.Join(cmdArgs, " "))
		fmt.Println(shellCmd.String())
		var out bytes.Buffer
		var stderr bytes.Buffer
		shellCmd.Stdout = &out
		shellCmd.Stderr = &stderr
		shellCmd.Env = os.Environ()

		err := shellCmd.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + out.String())
			return
		}
	},
}
