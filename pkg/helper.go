package pdf

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
)

func createTempFileName(tmpDir, ext string) string {
	return tmpDir + "/" + uuid.New().String() + "." + ext
}

func copyFile(src, dst string) error {

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}
	return nil
}

func prepareTempFilesForAppend(a *Append) error {
	a.TempOutput = createTempFileName(a.TempDir, "pdf")

	for _, input := range a.OriginalInputs {
		tempFileName := createTempFileName(a.TempDir, "pdf")

		if strings.HasSuffix(input, ".jpeg") || strings.HasSuffix(input, ".jpg") || strings.HasSuffix(input, ".png") {
			splittedFileName := strings.Split(input, ".")
			imageTempFileName := createTempFileName(a.TempDir, splittedFileName[len(splittedFileName)-1])
			err := copyFile(input, imageTempFileName)
			if err != nil {
				return err
			}

			err = Convert(imageTempFileName, tempFileName)
			if err != nil {
				return err
			}
			a.TempInputImages = append(a.TempInputImages, imageTempFileName)
			a.TempInputs = append(a.TempInputs, tempFileName)
			continue
		}

		err := copyFile(input, tempFileName)
		if err != nil {
			return fmt.Errorf("failed to copy file: %v", err)
		}
		a.TempInputs = append(a.TempInputs, tempFileName)
	}
	return nil
}
