package pdf

import "fmt"

// Stream is an onject
type Stream struct { // Object
	ObjectNumber int
	Data         []byte
	Filter       Filter
	Width        int
	Height       int
}

func (s Stream) ToString() string {
	var encodedData []byte
	switch s.Filter {
	case FlateDecode:
	default:
		encodedData = s.Data
	}

	return fmt.Sprintf("%d 0 obj\n<< /Length %d >>\nstream\n%s\nendstream\nendobj\n", s.ObjectNumber, len(encodedData), string(encodedData))
}

func (s Stream) GetObjectNumber() int {
	return s.ObjectNumber
}
