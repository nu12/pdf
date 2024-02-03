/*
Copyright © 2024 nu12
*/

package helper

import (
	"strings"

	"github.com/nu12/pdf/internal/model"
)

func UnmarshalImages(fileNames []string) []model.Imager {
	var files = make([]model.Imager, 0, len(fileNames))
	for _, fileName := range fileNames {
		splitted := strings.Split(fileName, ".")
		extention := splitted[len(splitted)-1]
		if strings.EqualFold("png", extention) {
			files = append(files, model.PNG{Filename: fileName})
			continue
		}
	}
	return files
}
