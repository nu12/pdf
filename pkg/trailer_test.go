package pdf

import (
	"testing"
)

func TestTrailerToString(t *testing.T) {

	trailer := Trailer{
		XRefTableStartOffset: 30,
		Size:                 4,
		Root:                 2,
		Info:                 3,
	}

	expected := "trailer\n<< /Size 4\n/Root 2 0 R\n/Info 3 0 R\n>>\nstartxref\n30\n%%EOF"
	result := trailer.ToString()

	if expected != result {
		t.Errorf("Expected Trailer to be:\n%s\nGot:\n%s\n", expected, result)
	}

}
