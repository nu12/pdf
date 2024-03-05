package pdf

import "fmt"

type Info struct {
	ObjectNumber int
	Title        string
	Producer     string
}

func (i Info) ToString() string {
	return fmt.Sprintf("%d 0 obj\n<</Title (%s)\n/Producer (%s)>>\nendobj\n", i.ObjectNumber, i.Title, i.Producer)
}

func (i Info) GetObjectNumber() int {
	return i.ObjectNumber
}
