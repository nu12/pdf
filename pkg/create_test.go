package pdf

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/nu12/pdf/internal/helper"
	"github.com/nu12/pdf/internal/model"
)

type CreateTestCase struct {
	Description        string
	Inputs             []model.Imager
	MarginProfile      model.MarginProfile
	CompressionProfile model.CompressionProfile
	ModeProfile        model.ModeProfile
	Error              error
}

func TestCreate(t *testing.T) {
	inputs := []string{"../examples/inputs/a4/1.png"}

	tests := []CreateTestCase{
		{Description: "Empty output", Inputs: []model.Imager{}, MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{Value: 1}, ModeProfile: model.ModeProfile{}, Error: nil},
		{Description: "Compression cannot be zero", Inputs: []model.Imager{}, MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{}, ModeProfile: model.ModeProfile{}, Error: errors.New("invalid compression option")},
		{Description: "File creation", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{Value: 32}, ModeProfile: model.ModeProfile{}, Error: nil},
		{Description: "With margins", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{X: 10, Y: 10}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{}, Error: nil},
		{Description: "A4P", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{Format: "A4", Orientation: "P", Width: 210, Height: 297}, Error: nil},
		{Description: "A4L", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{X: 40, Y: 40}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{Format: "A4", Orientation: "L", Width: 297, Height: 210}, Error: nil},
		{Description: "Wrong format", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{X: 40, Y: 40}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{Format: "A4", Orientation: "P"}, Error: nil},
		{Description: "Wrong orientation", Inputs: helper.UnmarshalImages(inputs), MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{Format: "x"}, Error: errors.New("unknown page size x")},
		{Description: "Input file doesn't exist", Inputs: helper.UnmarshalImages([]string{"no/such/file.png"}), MarginProfile: model.MarginProfile{}, CompressionProfile: model.CompressionProfile{Value: 128}, ModeProfile: model.ModeProfile{}, Error: errors.New("open no/such/file.png: no such file or directory")},
	}

	for _, tt := range tests {

		randOutput := fmt.Sprintf("../examples/outputs/%s.pdf", uuid.New().String())

		err := Create(tt.Inputs, randOutput, tt.MarginProfile, tt.CompressionProfile, tt.ModeProfile)
		if err != nil && (err.Error() != tt.Error.Error()) {
			t.Errorf("%s: expected error to be %v, got %v", tt.Description, tt.Error, err)
		}

		if _, fileErr := os.Stat(randOutput); err == nil && errors.Is(fileErr, os.ErrNotExist) {
			t.Errorf("%s: expected file %s to exist, but it doesn't", tt.Description, randOutput)
		}
	}
}
