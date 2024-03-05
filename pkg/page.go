package pdf

import "fmt"

type Page struct {
	ObjectNumber     int
	ContentReference int
	ParentReference  int
	ProcSetReference int
	Width            int
	Height           int
}

func (p Page) ToString() string {
	return fmt.Sprintf("%d 0 obj\n<< /Type /Page\n/Parent %d 0 R\n/MediaBox [ 0 0 %d %d ]\n/Contents %d 0 R\n/Resources << /ProcSet %d 0 R >>\n>>\nendobj\n", p.ObjectNumber, p.ParentReference, p.Width, p.Height, p.ContentReference, p.ProcSetReference)
}

func (p Page) GetObjectNumber() int {
	return p.ObjectNumber
}
