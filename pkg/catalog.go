package pdf

import "fmt"

type Catalog struct {
	ObjectNumber int
	OutlinesRef  int
	PagesRef     int
}

func (c Catalog) ToString() string {
	return fmt.Sprintf("%d 0 obj\n<< /Type /Catalog\n/Outlines %d 0 R\n/Pages %d 0 R\n>>\nendobj\n", c.ObjectNumber, c.OutlinesRef, c.PagesRef)
}

func (c Catalog) GetObjectNumber() int {
	return c.ObjectNumber
}
