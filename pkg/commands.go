package pdf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GhostScriptForAppend(a *Append) *Append {
	if _, err := exec.LookPath("gs"); err != nil {
		a.Err = fmt.Errorf("gs executable not found")
		return a
	}

	cmdArgs := []string{
		"gs",
		"-q",
		"-dNOPAUSE",
		"-dBATCH",
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=/ebook", // default
		"-dCompressFonts=true",
		"-sOutputFile=" + a.TempOutput,
		strings.Join(a.TempInputs, " "),
	}

	shellCmd := exec.Command("/bin/bash", "-c", strings.Join(cmdArgs, " "))
	a.Command = shellCmd.String()
	var out bytes.Buffer
	var stderr bytes.Buffer
	shellCmd.Stdout = &out
	shellCmd.Stderr = &stderr
	shellCmd.Env = os.Environ()

	err := shellCmd.Run()
	if err != nil {
		a.Err = fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
		return a
	}
	return a
}

func GhostScriptForSplit(s Split) (string, error) {
	return "", nil
}

func Convert(src, dst string) error {
	if _, err := exec.LookPath("convert"); err != nil {
		return fmt.Errorf("convert executable not found")
	}

	cmdArgs := []string{
		"convert",
		src,
		"-strip",
		dst,
	}

	shellCmd := exec.Command("/bin/bash", "-c", strings.Join(cmdArgs, " "))
	var out bytes.Buffer
	var stderr bytes.Buffer
	shellCmd.Stdout = &out
	shellCmd.Stderr = &stderr
	shellCmd.Env = os.Environ()

	err := shellCmd.Run()
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
	}
	return nil
}
