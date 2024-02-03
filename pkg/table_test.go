package pdf

import (
	"testing"
)

func TestXRefTableToString(t *testing.T) {

	table := XRefTable{Refs: []XRef{
		{Offset: 0, Generation: 65353, Type: "f"},
		{Offset: 15, Generation: 0, Type: "n"},
	}}

	expected := "xref\n0 2\n0000000000 65353 f \n0000000015 00000 n \n"
	result := table.ToString()
	if expected != result {
		t.Errorf("Expected XRefTable to be:\n%s\nGot:\n%s\n", expected, result)
	}
}
