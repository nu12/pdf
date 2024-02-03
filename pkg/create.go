/*
Copyright © 2024 nu12
*/

package pdf

import (
	"errors"

	"github.com/jung-kurt/gofpdf"
	"github.com/nu12/pdf/internal/model"
)

func Create(inputs []model.Imager, output string, mp model.MarginProfile, cp model.CompressionProfile, mode model.ModeProfile) error {
	if cp.Value == 0 {
		return errors.New("Invalid compression option.")
	}

	pdf := gofpdf.New(mode.Orientation, "mm", mode.Format, "")

	for _, img := range inputs {

		pdf.AddPage()

		rdr, err := img.Compress(cp)
		if err != nil {
			return err
		}

		pdf.RegisterImageOptionsReader(img.GetFilename(), gofpdf.ImageOptions{ImageType: img.GetType()}, rdr)
		pdf.ImageOptions(img.GetFilename(), float64(mp.X), float64(mp.Y), float64(mode.Width-(2*mp.X)), 0, false, gofpdf.ImageOptions{ImageType: img.GetType()}, 0, "")
	}

	return pdf.OutputFileAndClose(output)

}
