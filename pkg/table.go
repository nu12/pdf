package pdf

import "fmt"

type XRef struct {
	Offset     int
	Generation int
	Type       string
}

type XRefTable struct {
	Refs []XRef
}

func (xref XRefTable) ToString() string {
	s := fmt.Sprintf("xref\n%d %d\n", 0, len(xref.Refs))
	for _, ref := range xref.Refs {
		s += fmt.Sprintf("%.10d %.5d %s \n", ref.Offset, ref.Generation, ref.Type)
	}
	return s
}
