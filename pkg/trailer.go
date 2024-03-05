package pdf

import "fmt"

type Trailer struct {
	XRefTableStartOffset int
	Size                 int
	Root                 int
	Info                 int
}

func (t Trailer) ToString() string {
	return fmt.Sprintf("trailer\n<< /Size %d\n/Root %d 0 R\n/Info %d 0 R\n>>\nstartxref\n%d\n%%%%EOF", t.Size, t.Root, t.Info, t.XRefTableStartOffset)
}
