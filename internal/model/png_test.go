package model

import (
	"errors"
	"testing"
)

func TestPNGGetFilenameAndGetType(t *testing.T) {
	expectedFilename := "test.png"
	expectedType := "png"
	f := PNG{Filename: expectedFilename}
	if fn := f.GetFilename(); fn != expectedFilename {
		t.Errorf("Expected Filename to be %s, got %s", expectedFilename, fn)
	}

	if typ := f.GetType(); typ != expectedType {
		t.Errorf("Expected type to be %s, got %s", expectedType, typ)
	}
}

type PNGTestCase struct {
	Description string
	Input       PNG
	Error       error
}

func TestPNGCompress(t *testing.T) {
	tests := []PNGTestCase{
		{Description: "File does not exist", Input: PNG{Filename: "does/not/exist.png"}, Error: errors.New("open does/not/exist.png: no such file or directory")},
		{Description: "File cannot be decoded (is not PNG)", Input: PNG{Filename: "../../main.go"}, Error: errors.New("png: invalid format: not a PNG file")},
		{Description: "File is PNG", Input: PNG{Filename: "../../examples/inputs/a4/1.png"}, Error: nil},
	}

	cp := CompressionProfile{Value: 128}

	for _, tt := range tests {
		r, err := tt.Input.Compress(cp)
		if err != nil && (err.Error() != tt.Error.Error()) {
			t.Errorf("%s, expected error %v, got %v", tt.Description, tt.Error, err)
		}

		if err != nil && r == nil {
			t.Errorf("%s, expected reader to be not nil", tt.Description)
		}
	}
}
