package pdf

import "fmt"

type Outlines struct {
	ObjectNumber int
	Count        int
}

func (o Outlines) ToString() string {
	return fmt.Sprintf("%d 0 obj\n<< /Type /Outlines\n/Count %d\n>>\nendobj\n", o.ObjectNumber, o.Count)
}

func (o Outlines) GetObjectNumber() int {
	return o.ObjectNumber
}
