package helper

import "testing"

type UnmarshalImagesTestCase struct {
	Description  string
	Inputs       []string
	OutputLength int
}

func TestUnmarshalImages(t *testing.T) {

	tests := []UnmarshalImagesTestCase{
		{Description: "One valid file (png only)", Inputs: []string{"../examples/inputs/a4/1.png"}, OutputLength: 1},
		{Description: "Two valid files (png only)", Inputs: []string{"../examples/inputs/a4/1.png", "../examples/inputs/a4/2.png"}, OutputLength: 2},
		{Description: "One invalid file", Inputs: []string{"../main.go"}, OutputLength: 0},
		{Description: "Mixed valid and invalid files (png only)", Inputs: []string{"../examples/inputs/a4/1.png", "../examples/inputs/a4/2.png", "../main.go"}, OutputLength: 2},
	}

	for _, tt := range tests {
		imagers := UnmarshalImages(tt.Inputs)
		if len(imagers) != tt.OutputLength {
			t.Errorf("%s : expected length to be %d, got %d", tt.Description, tt.OutputLength, len(imagers))
		}
	}

}
