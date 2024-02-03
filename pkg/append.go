/*
Copyright © 2024 nu12
*/

package pdf

import (
	"fmt"
	"os"
	"slices"
)

type Append struct {
	OriginalInputs  []string
	OriginalOutput  string
	TempDir         string
	TempInputs      []string
	TempInputImages []string
	TempOutput      string
	Command         string
	Err             error
}

func InitAppend(args []string, tmpDir string) *Append {
	a := &Append{
		TempDir: tmpDir, // TODO: Remove trailing / if present
		Err:     nil,
	}

	if len(args) < 2 {
		a.Err = fmt.Errorf("you must provide at least two arguments: input file (or files) and an output file")
		return a
	}

	a.OriginalOutput = args[len(args)-1]
	a.OriginalInputs = args[:len(args)-1]

	return a
}

func (a *Append) PreProcess() *Append {
	if a.Err != nil {
		return a
	}

	a.Err = prepareTempFilesForAppend(a)
	return a
}

func (a *Append) Run() *Append {
	if a.Err != nil {
		return a
	}

	return GhostScriptForAppend(a)
}

func (a *Append) PostProcess() {
	exitCode := 0
	if a.Err != nil {
		fmt.Println(a.Err)
		exitCode = 1
	}

	err := copyFile(a.TempOutput, a.OriginalOutput)
	if err != nil {
		_ = fmt.Errorf("failed to create output file: %v", err)
		exitCode = 1
	}

	for _, file := range slices.Concat(append(a.TempInputs, a.TempOutput), a.TempInputImages) {
		err = os.Remove(file)
		if err != nil {
			_ = fmt.Errorf("failed to delete temp file: %v", err)
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}
