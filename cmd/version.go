/*
Copyright © 2024 nu12
*/

package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionShort bool
var version string = "v0.1.1"

func init() {
	versionCmd.Flags().BoolVarP(&versionShort, "short", "s", false, "Show short version")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long: `Show current version
	
Examples:
# Show current version of the tool and its dependencies:
pdf version

# Show current version of the tool only (short):
pdf version --short/-s`,

	Run: func(cmd *cobra.Command, args []string) {
		if versionShort {

			printVersion(versionShort)
		} else {
			printGhostscriptVersion()
			printImageMagickVersion()
			printVersion(versionShort)
			printNotice()
		}

	},
}

func printGhostscriptVersion() {
	gsVersion, err := exec.Command("gs", "--version").Output()
	if err != nil {
		fmt.Println("Ghostscript not found")
		return
	}
	fmt.Println("===========================")
	fmt.Println("=== Ghostscript version ===")
	fmt.Println("===========================")
	fmt.Println(string(gsVersion))
}

func printImageMagickVersion() {
	convertVersion, err := exec.Command("convert", "--version").Output()
	if err != nil {
		fmt.Println("ImageMagick not found")
		return
	}
	fmt.Println("===========================")
	fmt.Println("=== ImageMagick version ===")
	fmt.Println("===========================")
	fmt.Println(string(convertVersion))
}

func printVersion(short bool) {
	if short {
		fmt.Println(version)
	} else {
		fmt.Println("===========================")
		fmt.Println("==== nu12/pdf version =====")
		fmt.Println("===========================")
		fmt.Println(version)
		fmt.Println()
	}
}

func printNotice() {
	fmt.Println("===========================")
	fmt.Println("====== Contributing =======")
	fmt.Println("===========================")
	fmt.Println("Contribute to version 1.0.0, see: https://github.com/nu12/pdf?tab=readme-ov-file#contributing")
}
