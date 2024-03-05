package pdf

import "fmt"

type Header struct {
	Major int
	Minor int
}

func (h Header) ToString() string {
	return fmt.Sprintf("%%PDF-%d.%d\n", h.Major, h.Minor)
}
func (h Header) GetObjectNumber() int {
	return 0
}
